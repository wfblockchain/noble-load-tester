// package main

// import (
// 	"os"

// 	"github.com/cosmos/cosmos-sdk/server"
// 	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

// 	"github.com/wfblockchain/noble-fiattokenfactory/simapp"
// 	"github.com/wfblockchain/noble-fiattokenfactory/simapp/simd/cmd"
// )

// func main() {
// 	rootCmd, _ := cmd.NewRootCmd()

// 	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
// 		switch e := err.(type) {
// 		case server.ErrorCode:
// 			os.Exit(e.Code)

// 		default:
// 			os.Exit(1)
// 		}
// 	}
// }

package main

import (
	"os"

	"cosmossdk.io/simapp"
	"cosmossdk.io/simapp/simd/cmd"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}