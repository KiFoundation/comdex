package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "liquidation"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey  = "mem_liquidation"
)

var (
	LockedVaultIdKey     = []byte{0x01}
	LockedVaultKeyPrefix = []byte{0x11}
)

func LockedVaultKey(id uint64) []byte {
	return append(LockedVaultKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
