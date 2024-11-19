package main

import (
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func main() {
	moduleName := "blockrewards" // The name of your module account
	moduleAddress := authtypes.NewModuleAddress(moduleName)
	fmt.Printf("Address for module '%s': %s\n", moduleName, moduleAddress.String())
}
