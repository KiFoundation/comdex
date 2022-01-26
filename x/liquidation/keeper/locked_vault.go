package keeper

import (
	"fmt"
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	"github.com/comdex-official/comdex/x/liquidation/types"
	vaulttypes "github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuftypes "github.com/gogo/protobuf/types"
)

func (k *Keeper) SetLockedVaultID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.IDKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)

	store.Set(key, value)
}

func (k *Keeper) GetLockedVaultID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.IDKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k *Keeper) SetLockedVault(ctx sdk.Context, vault types.LockedVault) {
	var (
		store = k.Store(ctx)
		key   = types.VaultKey(vault.LockedVaultID)
		value = k.cdc.MustMarshal(&vault)
	)
	store.Set(key, value)
}

func (k *Keeper) GetLockedVault(ctx sdk.Context, id uint64) (vault types.LockedVault, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.VaultKey(id)
		value = store.Get(key)
	)
	if value == nil {
		return vault, false
	}
	_ = k.cdc.Unmarshal(value, &vault)
	return vault, true
}

func (k *Keeper) DeleteLockedVault(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.VaultKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) GetLockedVaults(ctx sdk.Context) (lockedvaults []types.LockedVault) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.VaultKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var lockedvault types.LockedVault
		k.cdc.MustUnmarshal(iter.Value(), &lockedvault)
		lockedvaults = append(lockedvaults, lockedvault)
	}

	return lockedvaults
}

func (k Keeper) getModAccountBalances(ctx sdk.Context, accountName string, denom string) sdk.Int {
	macc := k.GetModuleAccount(ctx, accountName)
	return k.GetBalance(ctx, macc.GetAddress(), denom).Amount
}

func (k Keeper) TransferCollateral(
	ctx sdk.Context,
	vault vaulttypes.Vault,
	assetIn assettypes.Asset,
	assetOut assettypes.Asset,
	collateralizationRatio sdk.Dec,
) error {

	collateralAvailableInVaultModule := k.getModAccountBalances(ctx, vaulttypes.ModuleName, assetIn.Denom)
	fmt.Println("{{{{{{{{{{{{{{{")
	fmt.Println(collateralAvailableInVaultModule)
	collateralCoins := sdk.NewCoin(assetIn.Denom, sdk.MinInt(vault.AmountIn, collateralAvailableInVaultModule))
	fmt.Println(collateralCoins)
	/*err := k.bank.SendCoinsFromModuleToModule(ctx, vaulttypes.ModuleName, types.ModuleName, sdk.NewCoins(collateralCoins))
	if err != nil {
		return err
	}
	cinliqmod := k.getModAccountBalances(ctx, types.ModuleName, assetIn.Denom)
	fmt.Println(cinliqmod)
	return nil*/
	return nil
}

/*func (k *Keeper) LockedVaultCollaterlizationRatio(
	ctx sdk.Context,
	amountIn sdk.Int,
	assetIn assettypes.Asset,
	amountOut sdk.Int,
	assetOut assettypes.Asset,
) (sdk.Dec, error) {

	assetInPrice, found := k.GetPriceForAsset(ctx, assetIn.Id)
	if !found {
		return sdk.ZeroDec(), types.ErrorPriceDoesNotExist
	}

	assetOutPrice, found := k.GetPriceForAsset(ctx, assetOut.Id)
	if !found {
		return sdk.ZeroDec(), types.ErrorPriceDoesNotExist
	}

	totalIn := amountIn.Mul(sdk.NewIntFromUint64(assetInPrice)).ToDec()
	if totalIn.LTE(sdk.ZeroDec()) {
		return sdk.ZeroDec(), types.ErrorInvalidAmountIn
	}

	totalOut := amountOut.Mul(sdk.NewIntFromUint64(assetOutPrice)).ToDec()
	if totalOut.LTE(sdk.ZeroDec()) {
		return sdk.ZeroDec(), types.ErrorInvalidAmountOut
	}

	return totalIn.Quo(totalOut), nil
}*/
