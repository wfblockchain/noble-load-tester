module github.com/kishanshukla-2307/noble-load-tester

go 1.22.1

require github.com/cosmos/cosmos-sdk v0.45.16

replace (

	// github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	// // github.com/wfblockchain/noblechain/v5 => /home/leo10/noblechain
	// github.com/cosmos/cosmos-sdk/simapp => cosmossdk.io/simapp v0.0.0-20240408150508-e5b0e0e4b245
	github.com/noble-assets/noble/v5/app => /home/leo10/noblechain
)
