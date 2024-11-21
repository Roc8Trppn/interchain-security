package blockrewards

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/math"
	keeper "github.com/Roc8Trppn/interchain-security/v6/x/ccv/provider/blockrewards/keeper"
	"github.com/Roc8Trppn/interchain-security/v6/x/ccv/provider/blockrewards/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ appmodule.HasEndBlocker = AppModule{}

// AppModuleBasic defines the basic application module used by the blockrewards module.
type AppModuleBasic struct{}

// Name returns the blockrewards module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the blockrewards module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// DefaultGenesis returns default genesis state as raw bytes for the blockrewards module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	genesis := types.DefaultGenesisState()
	return cdc.MustMarshalJSON(&genesis)
}

// ValidateGenesis validates the genesis state for the blockrewards module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, bz json.RawMessage) error {
	var genesisState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genesisState); err != nil {
		return err
	}
	return genesisState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the blockrewards module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// GetTxCmd returns the root tx command for the blockrewards module.
func (AppModuleBasic) GetTxCmd() *cobra.Command { return nil }

// GetQueryCmd returns the root query command for the blockrewards module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }

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

// IsAppModule is a marker method to identify AppModules
func (AppModule) IsAppModule() {}
func (AppModule) IsOnePerModuleType() {}
// RegisterInvariants registers the invariants for the blockrewards module.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// RegisterInterfaces registers the module's protobuf interfaces.
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// QuerierRoute returns the blockrewards module's query routing key.
func (AppModule) QuerierRoute() string { return types.ModuleName }

// RegisterServices registers the module's services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
}

// InitGenesis initializes the genesis state for the blockrewards module.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

// ExportGenesis exports the genesis state for the blockrewards module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genesis := types.DefaultGenesisState()
	return cdc.MustMarshalJSON(&genesis)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock executes all logic for the blockrewards module at the beginning of a block.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestFinalizeBlock) {}

// EndBlock executes all logic for the blockrewards module at the end of a block.
func (am AppModule) EndBlock(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.Logger().Info("Executing EndBlock for blockrewards module")
	EndBlocker(sdkCtx, ctx, am.keeper)
	return nil
}
// func (am AppModule) EndBlock(context Conte) (error) {
// 	ctx.Logger().Info("Entering EndBlock for blockrewards module")
// 	EndBlocker(ctx, am.keeper)
// 	return nil
// }
// EndBlocker is the core logic for the blockrewards module at the end of each block.
func EndBlocker (sdkContext sdk.Context, ctx context.Context, k keeper.Keeper) {
	// Define the block reward amount, e.g., 10000 stake
	
	if sdkContext.TxBytes() == nil || len(sdkContext.TxBytes()) == 0 {
        sdkContext.Logger().Info("Empty block detected, no transactions included.")
        return // Skip rewards for empty blocks
    }

	// Alternatively, check transaction count (if using a block header)
	newCoin := sdk.NewCoin("stake", math.NewInt(100000))
	rewardAmount := sdk.NewCoins(newCoin)

	// Call the reward distribution logic from the Keeper
	err := k.DistributeRewards(sdkContext, ctx, rewardAmount)
	if err != nil {
		sdkContext.Logger().Error("error in EndBlocker ", err.Error())
	}
}
