package keeper

import (
	"github.com/comdex-official/comdex/x/liquidation/types"
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

func (k *Keeper) SetVault(ctx sdk.Context, vault types.LockedVault) {
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
