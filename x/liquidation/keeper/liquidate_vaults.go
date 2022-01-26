package keeper

import (
	"time"

	assettypes "github.com/comdex-official/comdex/x/asset/types"
	liquidationtypes "github.com/comdex-official/comdex/x/liquidation/types"
	vaulttypes "github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuftypes "github.com/gogo/protobuf/types"
)

func (k Keeper) LiquidateVaults(ctx sdk.Context) error {
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
			err := k.TransferCollateralCreateLockedVaultAndDeleteVault(ctx, vault, assetIn, assetOut, collateralizationRatio)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (k Keeper) TransferCollateralCreateLockedVaultAndDeleteVault(
	ctx sdk.Context,
	vault vaulttypes.Vault,
	assetIn assettypes.Asset,
	assetOut assettypes.Asset,
	collateralizationRatio sdk.Dec,
) error {
	collateralAvailableInVaultModule := k.GetModAccountBalances(ctx, vaulttypes.ModuleName, assetIn.Denom)
	collateralCoins := sdk.NewCoin(assetIn.Denom, sdk.MinInt(vault.AmountIn, collateralAvailableInVaultModule))

	err := k.bank.SendCoinsFromModuleToModule(ctx, vaulttypes.ModuleName, liquidationtypes.ModuleName, sdk.NewCoins(collateralCoins))
	if err != nil {
		return err
	}

	locked_vault := liquidationtypes.LockedVault{
		Id:                   k.GetLockedVaultID(ctx) + 1,
		VaultId:              vault.ID,
		PairId:               vault.PairID,
		Owner:                vault.Owner,
		AmountIn:             vault.AmountIn,
		AmountOut:            vault.AmountOut,
		Initiator:            liquidationtypes.ModuleName,
		IsAuctioned:          false,
		CrAtLiquidation:      collateralizationRatio,
		LiquidationTimestamp: time.Now().UTC(),
	}
	k.SetLockedVaultID(ctx, locked_vault.Id)
	k.SetLockedVault(ctx, locked_vault)

	vaultOwner, _ := sdk.AccAddressFromBech32(vault.Owner)
	k.DeleteVault(ctx, vault.ID)
	k.DeleteVaultForAddressByPair(ctx, vaultOwner, vault.PairID)
	return nil
}

func (k Keeper) GetModAccountBalances(ctx sdk.Context, accountName string, denom string) sdk.Int {
	macc := k.GetModuleAccount(ctx, accountName)
	return k.GetBalance(ctx, macc.GetAddress(), denom).Amount
}

func (k *Keeper) GetLockedVaultID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = liquidationtypes.LockedVaultIdKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k *Keeper) SetLockedVaultID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = liquidationtypes.LockedVaultIdKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)
	store.Set(key, value)
}

func (k *Keeper) SetLockedVault(ctx sdk.Context, locked_vault liquidationtypes.LockedVault) {
	var (
		store = k.Store(ctx)
		key   = liquidationtypes.LockedVaultKey(locked_vault.Id)
		value = k.cdc.MustMarshal(&locked_vault)
	)
	store.Set(key, value)
}

func (k *Keeper) GetLockedVaults(ctx sdk.Context) (locked_vaults []liquidationtypes.LockedVault) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, liquidationtypes.LockedVaultKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var auction liquidationtypes.LockedVault
		k.cdc.MustUnmarshal(iter.Value(), &auction)
		locked_vaults = append(locked_vaults, auction)
	}

	return locked_vaults
}
