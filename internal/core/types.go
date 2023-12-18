package core

import rarimocore "github.com/rarimo/rarimo-core/x/rarimocore/types"

type IdentityStateTransferDetails struct {
	Operation *rarimocore.WorldCoinIdentityTransfer
	Proof     []byte
}
