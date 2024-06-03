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
package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wfblockchain/noble-cctp/x/cctp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UsedNonce(c context.Context, req *types.QueryGetUsedNonceRequest) (*types.QueryGetUsedNonceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	nonce := types.Nonce{
		SourceDomain: req.SourceDomain,
		Nonce:        req.Nonce,
	}
	found := k.GetUsedNonce(ctx, nonce)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUsedNonceResponse{Nonce: nonce}, nil
}

func (k Keeper) UsedNonces(c context.Context, req *types.QueryAllUsedNoncesRequest) (*types.QueryAllUsedNoncesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var usedNonces []types.Nonce
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	usedNonceStore := prefix.NewStore(store, types.KeyPrefix(types.UsedNonceKeyPrefix))

	pageRes, err := query.Paginate(usedNonceStore, req.Pagination, func(key []byte, value []byte) error {
		var usedNonce types.Nonce
		if err := k.cdc.Unmarshal(value, &usedNonce); err != nil {
			return err
		}

		usedNonces = append(usedNonces, usedNonce)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUsedNoncesResponse{UsedNonces: usedNonces, Pagination: pageRes}, nil
}
