package liquidation

import (
	"fmt"
	"github.com/comdex-official/comdex/x/liquidation/keeper"
	"github.com/comdex-official/comdex/x/liquidation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker compounds the debt in outstanding cdps and liquidates cdps that are below the required collateralization ratio
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
				vault = types.LockedVault{
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
			k.SetVault(ctx, vault)
		}
	}
}
