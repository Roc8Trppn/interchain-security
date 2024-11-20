package blockrewards

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"    //
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"       // For bank operations
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper" // For staking operations
)

// Keeper defines the blockrewards module's keeper
type Keeper struct {
    bankKeeper    bankKeeper.Keeper
    stakingKeeper stakingKeeper.Keeper
    accountKeeper accountKeeper.AccountKeeper
}

// NewKeeper creates a new blockrewards Keeper instance
func NewKeeper(
    bankKeeper bankKeeper.Keeper,
    stakingKeeper stakingKeeper.Keeper,
    accountKeeper accountKeeper.AccountKeeper,
) Keeper {
    return Keeper{
        bankKeeper:    bankKeeper,
        stakingKeeper: stakingKeeper,
        accountKeeper: accountKeeper,
    }
}

func (k Keeper) DistributeRewards(ctx sdk.Context, rewardAmount sdk.Coins) error {
	// Get all validators

	proposerAddress := ctx.BlockHeader().ProposerAddress
	proposerValidator, _ := k.stakingKeeper.ValidatorByConsAddr(ctx, sdk.ConsAddress(proposerAddress))
	if proposerValidator == nil {
        return fmt.Errorf("block proposer not found")
    }

	proposerOperatorAddress := proposerValidator.GetOperator()
	proposerAccAddress, err := sdk.AccAddressFromBech32(proposerOperatorAddress)
    if err != nil {
        return fmt.Errorf("invalid proposer address: %w", err)
    }
	
    err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "blockrewards", proposerAccAddress, rewardAmount)
	if err != nil {
		return fmt.Errorf("failed to send block rewards: %w", err)
	}

	ctx.Logger().Info("Distributed block reward", "proposer", proposerAccAddress.String(), "amount", rewardAmount.String())

	return nil
}
