package main

import (
	"github.com/informalsystems/tm-load-test/pkg/loadtest"
	"github.com/wfblockchain/noble-load-tester/pkg/noble"
)

func main() {
	if err := loadtest.RegisterClientFactory("noble", &noble.NobleClientFactory{}); err != nil {
		panic(err)
	}
	// The loadtest.Run method will handle CLI argument parsing, errors,
	// configuration, instantiating the load test and/or coordinator/worker
	// operations, etc. All it needs is to know which client factory to use for
	// its load testing.
	loadtest.Run(&loadtest.CLIConfig{
		AppName:              "noble-load-tester",
		AppShortDesc:         "Load testing application for noble (TM)",
		AppLongDesc:          "Some long description on how to use the tool",
		DefaultClientFactory: "kvstore",
	})
}
