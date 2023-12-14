package relayer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/worldcoin-relayer-svc/internal/config"
	"github.com/rarimo/worldcoin-relayer-svc/internal/core"
	"github.com/rarimo/worldcoin-relayer-svc/internal/data"
	"github.com/rarimo/worldcoin-relayer-svc/internal/data/pg"
	"github.com/rarimo/worldcoin-relayer-svc/internal/utils"
	"github.com/rarimo/worldcoin-relayer-svc/pkg/polygonid/contracts"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	ErrChainNotFound    = errors.New("chain not found")
	ErrEntryNotFound    = errors.New("entry not found")
	ErrAlreadySubmitted = errors.New("already transited")
)

type Service struct {
	log     *logan.Entry
	core    *core.Core
	chains  *config.EVM
	storage *pg.Storage
}

func NewService(cfg config.Config) *Service {
	return &Service{
		log:     cfg.Log(),
		core:    core.NewCore(cfg),
		chains:  cfg.EVM(),
		storage: pg.New(cfg.DB()),
	}
}

func (c *Service) StateRelays(ctx context.Context, state string) ([]data.Transition, error) {
	entry, err := c.storage.StateQ().StateByIDCtx(ctx, state, false)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get entry by state")
	}

	if entry == nil {
		return nil, ErrEntryNotFound
	}

	transitions, err := c.storage.TransitionQ().TransitionsByStateCtx(ctx, state, false)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transition")
	}

	return transitions, nil
}

func (c *Service) StateRelay(ctx context.Context, state string, chainName string, waitTxConfirm bool) (txhash string, err error) {
	chain, ok := c.chains.GetChainByName(chainName)
	if !ok {
		return "", ErrChainNotFound
	}

	entry, err := c.storage.StateQ().StateByIDCtx(ctx, state, false)
	if err != nil {
		return "", errors.Wrap(err, "failed to get entry by state")
	}

	if entry == nil {
		return "", ErrEntryNotFound
	}

	if err := c.checkTransitionNotExist(ctx, state, chainName); err != nil {
		return "", err
	}

	return c.processIdentityStateTransfer(ctx, chain, entry, waitTxConfirm)
}

func (c *Service) checkTransitionNotExist(ctx context.Context, state, chain string) error {
	transitions, err := c.storage.TransitionQ().TransitionsByStateCtx(ctx, state, false)
	if err != nil {
		return errors.Wrap(err, "failed to get transition")
	}

	if len(transitions) == 0 {
		return nil
	}

	for _, transition := range transitions {
		if transition.Chain == chain {
			return ErrAlreadySubmitted
		}
	}

	return nil
}

func (c *Service) processIdentityStateTransfer(ctx context.Context, chain *config.EVMChain, entry *data.State, waitTxConfirm bool) (txhash string, err error) {
	opts := chain.TransactorOpts()

	nonce, err := chain.RPC.PendingNonceAt(context.TODO(), chain.SubmitterAddress)
	if err != nil {
		return "", errors.Wrap(err, "failed to fetch a nonce")
	}

	opts.Nonce = big.NewInt(int64(nonce))

	opts.GasPrice, err = chain.RPC.SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to get suggested gas price")
	}

	details, err := c.core.GetIdentityStateTransferProof(ctx, entry.Operation)
	if err != nil {
		return "", errors.Wrap(err, "failed to get operation proof details")
	}

	contract, err := contracts.NewIdentityManager(chain.ContractAddress, chain.RPC)
	if err != nil {
		return "", errors.Wrap(err, "failed to create contract instance")
	}

	prevRoot := new(big.Int).SetBytes(hexutil.MustDecode(details.Operation.PrevState))
	root := new(big.Int).SetBytes(hexutil.MustDecode(details.Operation.State))
	timestamp, _ := new(big.Int).SetString(details.Operation.Timestamp, 10)

	tx, err := contract.SignedTransitRoot(opts, prevRoot, root, timestamp, details.Proof)
	if err != nil {
		c.log.Debugf(
			"Tx args: %s, %s, %s, %s",
			prevRoot.String(),
			root.String(),
			timestamp.String(),
			hexutil.Encode(details.Proof),
		)
		return "", errors.Wrap(err, "failed to send state transition tx")
	}

	transition := data.Transition{
		Tx:    tx.Hash().String(),
		State: entry.ID,
		Chain: chain.Name,
	}

	if err := c.storage.TransitionQ().Insert(&transition); err != nil {
		c.log.WithError(err).Error("failed to create transition entry")
	}

	if waitTxConfirm {
		c.waitTxConfirmation(ctx, chain, tx)
	}

	return tx.Hash().Hex(), nil
}

func (c *Service) waitTxConfirmation(ctx context.Context, chain *config.EVMChain, tx *types.Transaction) {
	receipt, err := bind.WaitMined(ctx, chain.RPC, tx)
	if err != nil {
		c.log.WithError(err).Error("failed to wait for state transition tx")
		return
	}

	if receipt.Status == 0 {
		c.log.WithError(err).WithFields(logan.F{
			"receipt": utils.Prettify(receipt),
			"chain":   chain.Name,
		}).Error("failed to wait for state transition tx")
		return
	}

	c.log.
		WithFields(logan.F{
			"tx_id":        tx.Hash(),
			"tx_index":     receipt.TransactionIndex,
			"block_number": receipt.BlockNumber,
			"gas_used":     receipt.GasUsed,
		}).
		Info("evm transaction confirmed")
}
