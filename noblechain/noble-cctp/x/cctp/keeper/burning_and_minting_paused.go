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
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/wfblockchain/noble-cctp/x/cctp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetBurningAndMintingPaused returns BurningAndMintingPaused
func (k Keeper) GetBurningAndMintingPaused(ctx sdk.Context) (val types.BurningAndMintingPaused, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurningAndMintingPausedKey))
	b := store.Get(types.KeyPrefix(types.BurningAndMintingPausedKey))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetBurningAndMintingPaused set BurningAndMintingPaused in the store
func (k Keeper) SetBurningAndMintingPaused(ctx sdk.Context, paused types.BurningAndMintingPaused) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurningAndMintingPausedKey))
	b := k.cdc.MustMarshal(&paused)
	store.Set(types.KeyPrefix(types.BurningAndMintingPausedKey), b)
}
