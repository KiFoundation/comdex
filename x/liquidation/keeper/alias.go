package keeper

import (
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	vaulttypes "github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) GetVaults(ctx sdk.Context) (vaults []vaulttypes.Vault) {
	return k.vault.GetVaults(ctx)
}

func (k *Keeper) DeleteVault(ctx sdk.Context, id uint64) {
	k.vault.DeleteVault(ctx, id)
}

func (k *Keeper) DeleteVaultForAddressByPair(ctx sdk.Context, address sdk.AccAddress, pairID uint64) {
	k.vault.DeleteVaultForAddressByPair(ctx, address, pairID)
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
