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
	"encoding/binary"
	"fmt"

	"github.com/wfblockchain/noble-cctp/x/cctp/types"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc              codec.BinaryCodec
		storeKey         storetypes.StoreKey
		paramstore       paramtypes.Subspace
		bank             types.BankKeeper
		fiattokenfactory types.FiatTokenfactoryKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bank types.BankKeeper,
	fiattokenfactory types.FiatTokenfactoryKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		paramstore:       ps,
		bank:             bank,
		fiattokenfactory: fiattokenfactory,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// func (k Keeper) setDone(ctx sdk.Context, name string) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set(encodeDoneKey(name, ctx.BlockHeight()), []byte{1})
// }

// encodeDoneKey - concatenate DoneByte, height and upgrade name to form the done key
func encodeDoneKey(name string, height int64) []byte {
	key := make([]byte, 9+len(name)) // 9 = donebyte + uint64 len
	key[0] = types.DoneByte
	binary.BigEndian.PutUint64(key[1:9], uint64(height))
	copy(key[9:], name)
	return key
}
