package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	appparams "github.com/Roc8Trppn/interchain-security/v6/app/params"
	app "github.com/Roc8Trppn/interchain-security/v6/app/sovereign"
	"github.com/Roc8Trppn/interchain-security/v6/cmd/interchain-security-sd/cmd"
)

func main() {
	appparams.SetAddressPrefixes(app.AccountAddressPrefix)
	rootCmd := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
