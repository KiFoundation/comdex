package expected

import (
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	liquidationtypes "github.com/comdex-official/comdex/x/liquidation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI
}

type BankKeeper interface {
	MintCoins(ctx sdk.Context, name string, coins sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type LiquidationKeeper interface {
	GetLockedVaults(ctx sdk.Context) (locked_vaults []liquidationtypes.LockedVault)
	FlagLockedVaultAsAuctioned(ctx sdk.Context, id uint64) error
	UnflagLockedVaultAsAuctioned(ctx sdk.Context, id uint64) error
}

type AssetKeeper interface {
	GetAsset(ctx sdk.Context, id uint64) (assettypes.Asset, bool)
	GetPair(ctx sdk.Context, id uint64) (assettypes.Pair, bool)
}

type VaultKeeper interface {
	CalculateCollaterlizationRatio(
		ctx sdk.Context,
		amountIn sdk.Int,
		assetIn assettypes.Asset,
		amountOut sdk.Int,
		assetOut assettypes.Asset,
	) (sdk.Dec, error)
}
