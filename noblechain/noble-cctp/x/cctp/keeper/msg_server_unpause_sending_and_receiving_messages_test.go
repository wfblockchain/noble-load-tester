/*
 * Copyright (c) 2023, © Circle Internet Financial, LTD.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
 * Authority not set
 * Invalid authority
 */
func TestUnpauseSendingAndReceivingMessagesHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	pauser := sample.AccAddress()
	testkeeper.SetPauser(ctx, pauser)

	message := types.MsgUnpauseSendingAndReceivingMessages{
		From: pauser,
	}
	_, err := server.UnpauseSendingAndReceivingMessages(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetSendingAndReceivingMessagesPaused(ctx)
	require.True(t, found)
	require.Equal(t, false, actual.Paused)
}

func TestUnpauseSendingAndReceivingMessagesAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgUnpauseSendingAndReceivingMessages{
		From: "authority",
	}

	require.PanicsWithValue(t, "cctp pauser not found in state", func() {
		_, _ = server.UnpauseSendingAndReceivingMessages(sdk.WrapSDKContext(ctx), &message)
	})
}

func TestUnpauseSendingAndReceivingMessagesInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	pauser := sample.AccAddress()
	testkeeper.SetPauser(ctx, pauser)

	message := types.MsgUnpauseSendingAndReceivingMessages{
		From: "not the authority",
	}
	_, err := server.UnpauseSendingAndReceivingMessages(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot unpause sending and receiving messages")
}
