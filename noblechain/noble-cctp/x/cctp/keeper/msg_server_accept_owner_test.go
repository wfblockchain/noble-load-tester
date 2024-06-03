package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/wfblockchain/noble-cctp/testutil/keeper"
	"github.com/wfblockchain/noble-cctp/testutil/sample"
	"github.com/wfblockchain/noble-cctp/x/cctp/keeper"
	"github.com/wfblockchain/noble-cctp/x/cctp/types"
)

/*
 * Happy path
 * Owner not set
 * Pending owner not set
 * Invalid Pending owner
 */

func TestAcceptOwnerHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)
	pendingOwner := sample.AccAddress()
	testkeeper.SetPendingOwner(ctx, pendingOwner)

	message := types.MsgAcceptOwner{
		From: pendingOwner,
	}

	_, err := server.AcceptOwner(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	newOwner := testkeeper.GetOwner(ctx)
	require.Equal(t, pendingOwner, newOwner)
}

func TestAcceptOwnerOwnerNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	pendingOwner := sample.AccAddress()
	testkeeper.SetPendingOwner(ctx, pendingOwner)

	message := types.MsgAcceptOwner{
		From: pendingOwner,
	}

	require.Panicsf(t, func() {
		_, _ = server.AcceptOwner(sdk.WrapSDKContext(ctx), &message)
	}, "cctp owner not found in state")
}

func TestAcceptOwnerPendingOwnerNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgAcceptOwner{
		From: sample.AccAddress(),
	}

	_, err := server.AcceptOwner(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "pending owner is not set")
}

func TestAcceptOwnerInvalidPendingOwner(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)
	pendingOwner := sample.AccAddress()
	testkeeper.SetPendingOwner(ctx, pendingOwner)

	message := types.MsgAcceptOwner{
		From: sample.AccAddress(),
	}

	_, err := server.AcceptOwner(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "you are not the pending owner")
}
