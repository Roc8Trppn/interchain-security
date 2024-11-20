package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	app "github.com/Roc8Trppn/interchain-security/v6/app/consumer-democracy"
	appparams "github.com/Roc8Trppn/interchain-security/v6/app/params"
	"github.com/Roc8Trppn/interchain-security/v6/cmd/interchain-security-cdd/cmd"
)

func main() {
	appparams.SetAddressPrefixes(app.Bech32MainPrefix)

	rootCmd := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
