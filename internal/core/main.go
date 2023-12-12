package core

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	rarimocore "github.com/rarimo/rarimo-core/x/rarimocore/types"
	"github.com/rarimo/worldcoin-relayer-svc/internal/config"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	proofType, _ = abi.NewType("bytes32[]", "", nil)
	sigType, _   = abi.NewType("bytes", "", nil)
	proofArgs    = abi.Arguments{
		{
			Name: "path",
			Type: proofType,
		},
		{
			Name: "signature",
			Type: sigType,
		},
	}
)

type Core struct {
	log  *logan.Entry
	core rarimocore.QueryClient
}

func NewCore(cfg config.Config) *Core {
	return &Core{
		core: rarimocore.NewQueryClient(cfg.Cosmos()),
		log:  cfg.Log().WithField("service", "core"),
	}
}

func (c *Core) GetIdentityStateTransferProof(ctx context.Context, operationID string) (*IdentityStateTransferDetails, error) {
	proof, err := c.core.OperationProof(ctx, &rarimocore.QueryGetOperationProofRequest{Index: operationID})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the operation proof")
	}

	pathHashes := make([]common.Hash, 0, len(proof.Path))
	for _, p := range proof.Path {
		pathHashes = append(pathHashes, common.HexToHash(p))
	}

	signature := hexutil.MustDecode(proof.Signature)
	signature[64] += 27

	operation, err := c.core.Operation(context.TODO(), &rarimocore.QueryGetOperationRequest{Index: operationID})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the operation")
	}

	transfer, err := pkg.GetWorldCoinIdentityTransfer(operation.Operation)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse operation details")
	}

	result := IdentityStateTransferDetails{Operation: transfer}

	result.Proof, err = proofArgs.Pack(pathHashes, signature)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode the proof")
	}

	return &result, nil
}
