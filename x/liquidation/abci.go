package liquidation

import (
	"fmt"

	"github.com/comdex-official/comdex/x/liquidation/keeper"
	"github.com/comdex-official/comdex/x/liquidation/types"
	vaulttypes "github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker compounds the debt in outstanding cdps and liquidates cdps that are below the required collateralization ratio
/*func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
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
}*/
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	vaults := k.GetVaults(ctx)
	for _, vault := range vaults {
		pair, found := k.GetPair(ctx, vault.PairID)
		if !found {
			continue
		}
		liquidationRatio := pair.LiquidationRatio
		assetIn, found := k.GetAsset(ctx, pair.AssetIn)
		if !found {
			continue
		}

		assetOut, found := k.GetAsset(ctx, pair.AssetOut)
		if !found {
			continue
		}
		collateralizationRatio, err := k.CalculateCollaterlizationRatio(ctx, vault.AmountIn, assetIn, vault.AmountOut, assetOut)
		if err != nil {
			continue
		}
		if sdk.Dec.LT(collateralizationRatio, liquidationRatio) {
			fmt.Println("Liquidating here")
			var (
				id    = k.GetLockedVaultID(ctx)
				LockedVault = types.LockedVault{
					LockedVaultID: id + 1,
					ID:            vault.ID,
					PairID:        vault.PairID,
					Owner:         vault.Owner,
					AmountIn:      vault.AmountIn,
					AmountOut:     vault.AmountOut,
					NewOwner:      types.ModuleName,
				}
			)
			k.SetLockedVaultID(ctx, id+1)
			k.SetLockedVault(ctx, LockedVault)
			balance1 := k.GetBalance(ctx, k.GetModuleAccount(ctx, types.ModuleName).GetAddress(), "ucmdx").Amount
			balance2 := k.GetBalance(ctx, k.GetModuleAccount(ctx, vaulttypes.ModuleName).GetAddress(), "ucmdx").Amount
			fmt.Println("liquidation : ", balance1)
			fmt.Println("vault : ", balance2)

			/*err := k.SendCoinsFromModuleToModule(ctx, vaulttypes.ModuleName, types.ModuleName, sdk.NewCoins(sdk.NewCoin("ucmdx", sdk.NewInt(10))))
			if err != nil {
				return
			}

			fmt.Println("After Transfer")*/
			k.DeleteVault(ctx, LockedVault.ID)
		}
	}

	lockedvaults := k.GetLockedVaults(ctx)

	for _, lockedvault := range lockedvaults {
		pair, found := k.GetPair(ctx, lockedvault.PairID)
		if !found {
			continue
		}
		liquidationRatio := pair.LiquidationRatio
		assetIn, found := k.GetAsset(ctx, pair.AssetIn)
		if !found {
			continue
		}

		assetOut, found := k.GetAsset(ctx, pair.AssetOut)
		if !found {
			continue
		}

		collateralizationRatio, err := k.CalculateCollaterlizationRatio(ctx, lockedvault.AmountIn, assetIn, lockedvault.AmountOut, assetOut)
		if err != nil {
			continue
		}
		if sdk.Dec.GT(collateralizationRatio,liquidationRatio){
			fmt.Println("UnLiquidating here")
			var (
				id    = k.GetID(ctx)
				Vault = vaulttypes.Vault{
					ID:            lockedvault.ID,
					PairID:        lockedvault.PairID,
					Owner:         lockedvault.Owner,
					AmountIn:      lockedvault.AmountIn,
					AmountOut:     lockedvault.AmountOut,
				}
			)
			k.SetID(ctx, id+1)
			k.SetVault(ctx, Vault)
			k.DeleteLockedVault(ctx, lockedvault.LockedVaultID)
		}
	}
}
