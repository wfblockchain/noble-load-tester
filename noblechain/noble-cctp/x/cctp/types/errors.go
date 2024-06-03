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

// DONTCOVER

import (
	errorof "cosmossdk.io/errors"
)

// x/cctp module sentinel errors
var (
	ErrUnauthorized                     = errorof.Register(ModuleName, 30, "unauthorized")
	ErrMint                             = errorof.Register(ModuleName, 31, "tokens can not be minted")
	ErrBurn                             = errorof.Register(ModuleName, 32, "tokens can not be burned")
	ErrAttesterAlreadyFound             = errorof.Register(ModuleName, 33, "attester is already present")
	ErrAuthorityNotSet                  = errorof.Register(ModuleName, 34, "authority not set")
	ErrMalformedField                   = errorof.Register(ModuleName, 35, "field cannot be empty or nil")
	ErrReceiveMessage                   = errorof.Register(ModuleName, 36, "err in receive message")
	ErrDisableAttester                  = errorof.Register(ModuleName, 37, "err in disable attester")
	ErrUpdateSignatureThreshold         = errorof.Register(ModuleName, 38, "err in update signature threshold")
	ErrMinterAllowanceNotFound          = errorof.Register(ModuleName, 39, "minter allowance not found")
	ErrTokenPairAlreadyFound            = errorof.Register(ModuleName, 40, "token pair already exists")
	ErrTokenPairNotFound                = errorof.Register(ModuleName, 41, "token pair not found")
	ErrSendMessage                      = errorof.Register(ModuleName, 42, "error in send message")
	ErrSendMessageWithCaller            = errorof.Register(ModuleName, 43, "error in send message with caller")
	ErrDepositForBurn                   = errorof.Register(ModuleName, 44, "error in deposit for burn")
	ErrInvalidDestinationCaller         = errorof.Register(ModuleName, 45, "malformed destination caller")
	ErrSignatureVerification            = errorof.Register(ModuleName, 46, "unable to verify signature")
	ErrReplaceMessage                   = errorof.Register(ModuleName, 47, "error in replace message")
	ErrDuringPause                      = errorof.Register(ModuleName, 48, "error while trying to pause or unpause")
	ErrInvalidAmount                    = errorof.Register(ModuleName, 49, "invalid amount")
	ErrNextAvailableNonce               = errorof.Register(ModuleName, 50, "error while retrieving next available nonce")
	ErrRemoteTokenMessengerAlreadyFound = errorof.Register(ModuleName, 51, "this remote token messenger mapping already exists")
	ErrRemoteTokenMessengerNotFound     = errorof.Register(ModuleName, 53, "remote token messenger not found")
	ErrParsingMessage                   = errorof.Register(ModuleName, 54, "error while parsing message into bytes")
	ErrParsingBurnMessage               = errorof.Register(ModuleName, 55, "error while parsing burn message into bytes")
	ErrInvalidRemoteToken               = errorof.Register(ModuleName, 56, "invalid remote token")
)
