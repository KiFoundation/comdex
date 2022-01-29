package keeper

import (
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	liquidationtypes "github.com/comdex-official/comdex/x/liquidation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k *Keeper) GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI {
	return k.account.GetModuleAccount(ctx, name)
}

func (k *Keeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return k.bank.GetBalance(ctx, addr, denom)
}

func (k *Keeper) MintCoin(ctx sdk.Context, name string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.bank.MintCoins(ctx, name, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error {
	return k.bank.SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt)
}

func (k *Keeper) GetLockedVaults(ctx sdk.Context) (locked_vaults []liquidationtypes.LockedVault) {
	return k.liquidation.GetLockedVaults(ctx)
}

func (k *Keeper) GetPair(ctx sdk.Context, id uint64) (assettypes.Pair, bool) {
	return k.asset.GetPair(ctx, id)
}

func (k *Keeper) GetAsset(ctx sdk.Context, id uint64) (assettypes.Asset, bool) {
	return k.asset.GetAsset(ctx, id)
}

func (k *Keeper) CalculateCollaterlizationRatio(
	ctx sdk.Context,
	amountIn sdk.Int,
	assetIn assettypes.Asset,
	amountOut sdk.Int,
	assetOut assettypes.Asset,
) (sdk.Dec, error) {
	return k.vault.CalculateCollaterlizationRatio(ctx, amountIn, assetIn, amountOut, assetOut)
}

func (k *Keeper) FlagLockedVaultAsAuctioned(ctx sdk.Context, id uint64) error {
	return k.liquidation.FlagLockedVaultAsAuctioned(ctx, id)
}

func (k *Keeper) UnflagLockedVaultAsAuctioned(ctx sdk.Context, id uint64) error {
	return k.liquidation.UnflagLockedVaultAsAuctioned(ctx, id)
}
