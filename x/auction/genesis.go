package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gauss/gauss/v6/x/auction/keeper"
	"github.com/gauss/gauss/v6/x/auction/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	orders := k.GetAllOrders(ctx)
	params := types.Params{AutoAgreePeriod: k.GetParams(ctx).AutoAgreePeriod, Orders: orders}
	return types.NewGenesis(params)
}
