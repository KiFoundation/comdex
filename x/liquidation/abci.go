package liquidation

import (
	"fmt"

	"github.com/comdex-official/comdex/x/liquidation/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	fmt.Println("Printing Locked Vaults")
	fmt.Println(k.GetLockedVaults(ctx))
	fmt.Println(k.GetLockedVault(ctx, 1))
	fmt.Println(k.GetLockedVault(ctx, 2))
	k.LiquidateVaults(ctx)
	k.UnLiquidateVaults(ctx)
}
