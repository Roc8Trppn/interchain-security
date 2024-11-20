package blockrewards

import (
	"cosmossdk.io/math"
	"github.com/Roc8Trppn/interchain-security/v6/x/ccv/provider/blockrewards/keeper"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AppModuleBasic defines the basic application module used by the blockrewards module.
type AppModuleBasic struct{}

// Name returns the blockrewards module's name.
func (AppModuleBasic) Name() string {
	return "blockrewards"
}

// AppModule implements the AppModule interface for the blockrewards module.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

// EndBlock is called at the end of each block and triggers the `EndBlocker` logic.
func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestFinalizeBlock) []abci.ValidatorUpdate {
	EndBlocker(ctx, am.keeper)
	return []abci.ValidatorUpdate{}
}

// EndBlocker is the core logic for the blockrewards module at the end of each block.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// Define the block reward amount, e.g., 10000 stake
	newCoin := sdk.NewCoin("stake", math.NewInt(100000))
	rewardAmount := sdk.NewCoins(newCoin)

	// Call the reward distribution logic from the Keeper
	err := k.DistributeRewards(ctx, rewardAmount)
	if err != nil {
		ctx.Logger().Error("Failed to distribute block rewards", "error", err)
	}
}
