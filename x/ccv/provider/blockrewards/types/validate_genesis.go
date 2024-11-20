package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate checks the validity of the GenesisState.
func (gs GenesisState) Validate() error {
    if err := sdk.Coins(gs.ModuleAccountBalance).Validate(); err != nil {
        return fmt.Errorf("invalid module account balance: %w", err)
    }
    return nil
}