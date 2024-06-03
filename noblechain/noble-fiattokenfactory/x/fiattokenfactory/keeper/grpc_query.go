package keeper

import (
	"github.com/wfblockchain/noble-fiattokenfactory/x/fiattokenfactory/types"
)

var _ types.QueryServer = Keeper{}
