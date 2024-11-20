package types

import (
	"cosmossdk.io/math"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ModuleName defines the name of the module
const ModuleName = "blockrewards"

// DefaultGenesisState returns the default genesis state for the blockrewards module.
func DefaultGenesisState() GenesisState {
    return GenesisState{
        ModuleAccountBalance: sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(0))),
    }
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// If you have interfaces to register, add them here. For example:
	// registry.RegisterImplementations(
	//     (*sdk.Msg)(nil), // Register any Msg types here if needed
	// )
}