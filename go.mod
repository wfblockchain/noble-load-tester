module github.com/wfblockchain/noble-load-tester

go 1.21

replace (
	github.com/ChainSafe/go-schnorrkel => github.com/ChainSafe/go-schnorrkel v0.0.0-20200405005733-88cbf1b4c40d
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	// github.com/informalsystems/tm-load-test => ../tm-load-test

	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.16.0
	github.com/prometheus/client_model => github.com/prometheus/client_model v0.3.0
	github.com/prometheus/common => github.com/prometheus/common v0.42.0
	github.com/prometheus/procfs => github.com/prometheus/procfs v0.10.1
	github.com/tendermint/tendermint => github.com/cometbft/cometbft v0.37.2
	// // github.com/wfblockchain/noblechain/v5 => /home/leo10/noblechain
	// github.com/cosmos/cosmos-sdk/simapp => cosmossdk.io/simapp v0.0.0-20240408150508-e5b0e0e4b245
	github.com/wfblockchain/noblechain/v5/app => ../noblechain/
	golang.org/x/exp => golang.org/x/exp v0.0.0-20230711153332-06a737ee72cb
)
