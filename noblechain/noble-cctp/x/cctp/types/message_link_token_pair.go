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
package types

import (
	errorof "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLinkTokenPair = "link_token_pair"

var _ sdk.Msg = &MsgLinkTokenPair{}

func NewMsgLinkTokenPair(from string, localToken string, remoteToken []byte, remoteDomain uint32) *MsgLinkTokenPair {
	return &MsgLinkTokenPair{
		From:         from,
		LocalToken:   localToken,
		RemoteToken:  remoteToken,
		RemoteDomain: remoteDomain,
	}
}

func (msg *MsgLinkTokenPair) Route() string {
	return RouterKey
}

func (msg *MsgLinkTokenPair) Type() string {
	return TypeMsgLinkTokenPair
}

func (msg *MsgLinkTokenPair) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgLinkTokenPair) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLinkTokenPair) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorof.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address: %s", err)
	}

	if len(msg.RemoteToken) != 32 {
		return errorof.Wrapf(ErrInvalidRemoteToken, "must be a byte32 array: %s", err)
	}

	return nil
}