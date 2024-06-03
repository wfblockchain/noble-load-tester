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

func (k Keeper) Attester(c context.Context, req *types.QueryGetAttesterRequest) (*types.QueryGetAttesterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAttester(ctx, req.Attester)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAttesterResponse{Attester: val}, nil
}

func (k Keeper) Attesters(c context.Context, req *types.QueryAllAttestersRequest) (*types.QueryAllAttestersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var attesters []types.Attester
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	attestersStore := prefix.NewStore(store, types.KeyPrefix(types.AttesterKeyPrefix))

	pageRes, err := query.Paginate(attestersStore, req.Pagination, func(key []byte, value []byte) error {
		var Attester types.Attester
		if err := k.cdc.Unmarshal(value, &Attester); err != nil {
			return err
		}

		attesters = append(attesters, Attester)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAttestersResponse{Attesters: attesters, Pagination: pageRes}, nil
}
