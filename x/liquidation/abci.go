package liquidation

import (
	"fmt"

	"github.com/comdex-official/comdex/x/liquidation/keeper"
	liquidationtypes "github.com/comdex-official/comdex/x/liquidation/types"
	vaulttypes "github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker compounds the debt in outstanding cdps and liquidates cdps that are below the required collateralization ratio
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	balance1 := k.GetBalance(ctx, k.GetModuleAccount(ctx, liquidationtypes.ModuleName).GetAddress(), "ucmdx").Amount
	balance2 := k.GetBalance(ctx, k.GetModuleAccount(ctx, vaulttypes.ModuleName).GetAddress(), "ucmdx").Amount
	fmt.Println("liquidation : ", balance1)
	fmt.Println("vault : ", balance2)

	k.SendCoinsFromModuleToModule(ctx, vaulttypes.ModuleName, liquidationtypes.ModuleName, sdk.NewCoins(sdk.NewCoin("ucmdx", sdk.NewInt(10))))

	fmt.Println("After Transfer")
	balance3 := k.GetBalance(ctx, k.GetModuleAccount(ctx, liquidationtypes.ModuleName).GetAddress(), "ucmdx").Amount
	balance4 := k.GetBalance(ctx, k.GetModuleAccount(ctx, vaulttypes.ModuleName).GetAddress(), "ucmdx").Amount
	fmt.Println("liquidation : ", balance3)
	fmt.Println("vault : ", balance4)
}
