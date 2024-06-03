package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/wfblockchain/noblechain/v5/testutil/keeper"
	"github.com/wfblockchain/noblechain/v5/testutil/nullify"
	"github.com/wfblockchain/noblechain/v5/x/tokenfactory/keeper"
	"github.com/wfblockchain/noblechain/v5/x/tokenfactory/types"
)

func createTestMintingDenom(keeper *keeper.Keeper, ctx sdk.Context) types.MintingDenom {
	item := types.MintingDenom{
		Denom: "abcd",
	}
	keeper.SetMintingDenom(ctx, item)
	return item
}

func TestMintingDenomGet(t *testing.T) {
	keeper, ctx := keepertest.TokenfactoryKeeper(t)
	item := createTestMintingDenom(keeper, ctx)
	rst := keeper.GetMintingDenom(ctx)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}