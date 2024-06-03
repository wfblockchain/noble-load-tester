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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/wfblockchain/noble-cctp/testutil/keeper"
	"github.com/wfblockchain/noble-cctp/testutil/nullify"
	"github.com/wfblockchain/noble-cctp/x/cctp/types"
)

func TestNextAvailableNonceQuery(t *testing.T) {
	Nonce := types.Nonce{Nonce: uint64(123)}

	for _, tc := range []struct {
		desc     string
		set      bool
		request  *types.QueryGetNextAvailableNonceRequest
		response *types.QueryGetNextAvailableNonceResponse
		err      error
	}{
		{
			desc:     "HappyPath",
			set:      true,
			request:  &types.QueryGetNextAvailableNonceRequest{},
			response: &types.QueryGetNextAvailableNonceResponse{Nonce: Nonce},
		},
		{
			desc:    "NotFound",
			set:     false,
			request: &types.QueryGetNextAvailableNonceRequest{},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := keepertest.CctpKeeper(t)
			goCtx := sdk.WrapSDKContext(ctx)

			if tc.set {
				keeper.SetNextAvailableNonce(ctx, Nonce)
			}

			response, err := keeper.NextAvailableNonce(goCtx, tc.request)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
