package blockrewards

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"
	keeper "github.com/Roc8Trppn/interchain-security/v6/x/ccv/provider/blockrewards/keeper"
	"github.com/Roc8Trppn/interchain-security/v6/x/ccv/provider/blockrewards/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
     _ module.HasABCIGenesis = (*AppModule)(nil)
	 _ appmodule.HasEndBlocker = AppModule{}
)

// AppModuleBasic defines the basic application module used by the blockrewards module.
type AppModuleBasic struct{}

// Name returns the blockrewards module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the blockrewards module's types on the LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the blockrewards module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

// GetTxCmd returns the root tx command for the blockrewards module.
func (AppModuleBasic) GetTxCmd() *cobra.Command { return nil }

// GetQueryCmd returns the root query command for the blockrewards module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }

// AppModule implements the AppModule interface for the blockrewards module.
type AppModule struct {
	cdc codec.Codec
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(cdc codec.Codec, k keeper.Keeper) AppModule {
	return AppModule{
		cdc: cdc,
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

// // DefaultGenesis returns default genesis state as raw bytes for the module.
// func (am AppModule) DefaultGenesis() json.RawMessage {
// 	return am.cdc.MustMarshalJSON(types.NewGenesisState())
// }
// func (am AppModule) DefaultGenesis(writerFn func(string) (io.WriteCloser, error)) error {
//    // Use the provided writer function to obtain the writer
	
//     writer, err := writerFn("default")
//     if err != nil {
//         return fmt.Errorf("failed to create writer for genesis data: %w", err)
//     }
//     defer writer.Close()

//     // Marshal the default genesis state into JSON
//     data, err := am.cdc.MarshalJSON(types.NewGenesisState())
//     if err != nil {
//         return fmt.Errorf("failed to marshal default genesis state: %w", err)
//     }

//     // Write the data to the writer
//     if _, err := writer.Write(data); err != nil {
//         return fmt.Errorf("failed to write default genesis data: %w", err)
//     }

//     return nil
// }


// ValidateGenesis performs genesis state validation for the circuit module.
// func (am AppModule) ValidateGenesis(bz json.RawMessage) error {
// 	var data types.GenesisState
// 	if err := am.cdc.UnmarshalJSON(bz, &data); err != nil {
// 		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
// 	}

// 	return data.Validate()
// }
// ValidateGenesis performs genesis state validation for the circuit module.
// func (am AppModule) ValidateGenesis(readerFn func(string) (io.ReadCloser, error)) error {
//     reader, err := readerFn("default")
//     if err != nil {
//         return err
//     }
//     defer reader.Close()

//     data, err := io.ReadAll(reader)
//     if err != nil {
//         return err
//     }

//     var genesisState types.GenesisState
//     if err := am.cdc.UnmarshalJSON(data, &genesisState); err != nil {
//         return fmt.Errorf("failed to unmarshal genesis state: %w", err)
//     }

//     return genesisState.Validate()
// }

// InitGenesis initializes the genesis state for the blockrewards module.
// func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) error {
// 	// Unwrap the sdk.Context from the standard context
// 	sdkCtx := sdk.UnwrapSDKContext(ctx)
// 	// Unmarshal the provided genesis data into the module's GenesisState
// 	var genesisState types.GenesisState

// 	if err := am.cdc.UnmarshalJSON(data, &genesisState); err != nil {
// 		return fmt.Errorf("failed to unmarshal blockrewards genesis state: %w", err)
// 	}

// 	if err := am.keeper.InitGenesis(ctx, &genesisState); err != nil {
// 		return fmt.Errorf("failed to initialize %s genesis state: %v", types.ModuleName, err)
// 	}

// 	sdkCtx.Logger().Info("Successfully initialized blockrewards module genesis state")
// 	return nil
// }
// func (am AppModule) InitGenesis(ctx context.Context, readerFn func(string) (io.ReadCloser, error)) error {
//     reader, err := readerFn("default")
//     if err != nil {
//         return err
//     }
//     defer reader.Close()

//     data, err := io.ReadAll(reader)
//     if err != nil {
//         return err
//     }

//     var genesisState types.GenesisState
//     if err := am.cdc.UnmarshalJSON(data, &genesisState); err != nil {
//         return fmt.Errorf("failed to unmarshal genesis state: %w", err)
//     }

//     return am.keeper.InitGenesis(sdk.UnwrapSDKContext(ctx), &genesisState)
// }

func (am AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
    // Return the default genesis state marshaled as JSON
    gen := types.DefaultGenesisState()
    return cdc.MustMarshalJSON(&gen)
}

func (am AppModule) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
    // Unmarshal the genesis state from the JSON
    var genState types.GenesisState
    if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
        return fmt.Errorf("failed to unmarshal genesis state: %w", err)
    }

    // Validate the genesis state
    return genState.Validate()
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
    var genesisState types.GenesisState
    cdc.MustUnmarshalJSON(data, &genesisState)
    am.keeper.InitGenesis(ctx, &genesisState)
    return nil
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
     gs := am.keeper.ExportGenesis(ctx)
    return cdc.MustMarshalJSON(gs)
}

// ExportGenesis exports the genesis state for the blockrewards module.
// func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
// 	gs, err := am.keeper.ExportGenesis(ctx)
// 	if err != nil {
// 		return nil
// 	}

// 	return am.cdc.MustMarshalJSON(gs)
// }
// func (am AppModule) ExportGenesis(ctx context.Context, writerFn func(string) (io.WriteCloser, error)) error {
//     gs, err := am.keeper.ExportGenesis(sdk.UnwrapSDKContext(ctx))
//     if err != nil {
//         return err
//     }

//     // Marshal the genesis state directly
//     data, err := am.cdc.MarshalJSON(gs)
//     if err != nil {
//         return err
//     }

//     writer, err := writerFn("")
//     if err != nil {
//         return err
//     }
//     defer writer.Close()

//     _, err = writer.Write(data)
//     return err
// }

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
	
	// if sdkContext.TxBytes() == nil || len(sdkContext.TxBytes()) == 0 {
    //     sdkContext.Logger().Info("Empty block detected, no transactions included.")
    //     return // Skip rewards for empty blocks
    // }
	params, err := k.GetParams(sdkContext)
    if err != nil {
        sdkContext.Logger().Error("failed to get block reward params in EndBlocker", "error", err.Error())
        return
    }

    // Use the block_reward_amount from the retrieved params
    rewardAmount := sdk.NewCoins(params.BlockRewardAmount)

	// Call the reward distribution logic from the Keeper
	err2 := k.DistributeRewards(sdkContext, ctx, rewardAmount)
	if err2 != nil {
		sdkContext.Logger().Error("error in EndBlocker ", err2.Error())
	}
}
