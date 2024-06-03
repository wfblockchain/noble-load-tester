package simapp

import (
	"github.com/cosmos/cosmos-sdk/std"

	// simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	// paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	// simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
	// simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	simappparams "cosmossdk.io/simapp/params"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing. This function
// should be used only in tests or when creating a new app instance (NewApp*()).
// App user shouldn't create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
func MakeTestEncodingConfig() simappparams.EncodingConfig {
	encodingConfig := simappparams.MakeTestEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
