package keeper

import (
	"fmt"
	"time"

	assettypes "github.com/comdex-official/comdex/x/asset/types"
	auctiontypes "github.com/comdex-official/comdex/x/auction/types"
	liquidationtypes "github.com/comdex-official/comdex/x/liquidation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuftypes "github.com/gogo/protobuf/types"
)

func (k Keeper) CreateNewAuctions(ctx sdk.Context) {
	locked_vaults := k.GetLockedVaults(ctx)
	for _, locked_vault := range locked_vaults {
		fmt.Println(locked_vault)
		pair, found := k.GetPair(ctx, locked_vault.PairId)
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
		collateralizationRatio, err := k.CalculateCollaterlizationRatio(ctx, locked_vault.AmountIn, assetIn, locked_vault.AmountOut, assetOut)
		if err != nil {
			continue
		}
		if sdk.Dec.LT(collateralizationRatio, liquidationRatio) && !locked_vault.IsAuctioned {
			k.StartCollateralAuction(ctx, locked_vault, assetIn, assetOut, liquidationRatio)
		}
	}
}

func (k Keeper) CloseAuctions(ctx sdk.Context) {
	collateral_auctions := k.GetCollateralAuctions(ctx)
	for _, collateral_auction := range collateral_auctions {
		if ctx.BlockTime().After(collateral_auction.EndTime) {
			k.CloseCollateralAuction(ctx, collateral_auction)
		}
	}
}

func (k Keeper) StartCollateralAuction(
	ctx sdk.Context,
	locked_vault liquidationtypes.LockedVault,
	assetIn assettypes.Asset,
	assetOut assettypes.Asset,
	liquidationRatio sdk.Dec,
) error {

	auction := auctiontypes.CollateralAuction{
		LockedVaultId: locked_vault.Id,
		Collateral:    sdk.NewCoin(assetIn.Denom, locked_vault.AmountIn),
		Debt:          sdk.NewCoin(assetOut.Denom, locked_vault.AmountOut),
		Bidder:        nil,
		Bid:           sdk.NewCoin(assetOut.Denom, sdk.NewInt(0)),
		MinBid:        sdk.NewCoin(assetOut.Denom, locked_vault.AmountOut),
		MaxBid:        sdk.NewCoin(assetOut.Denom, locked_vault.AmountOut),
		EndTime:       time.Now().Local().Add(time.Hour * time.Duration(24)),
	}
	auction.Id = k.GetCollateralAuctionID(ctx) + 1
	k.SetCollateralAuctionID(ctx, auction.Id)
	k.SetCollateralAuction(ctx, auction)
	k.FlagLockedVaultAsAuctioned(ctx, locked_vault.Id)
	return nil
}

func (k Keeper) CloseCollateralAuction(
	ctx sdk.Context,
	collateral_auction auctiontypes.CollateralAuction,
) {
	k.UnflagLockedVaultAsAuctioned(ctx, collateral_auction.LockedVaultId)
	k.DeleteCollateralAuction(ctx, collateral_auction.Id)
}

func (k *Keeper) GetCollateralAuctionID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionIdKey
		value = store.Get(key)
	)
	if value == nil {
		return 0
	}
	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k *Keeper) SetCollateralAuctionID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionIdKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)

	store.Set(key, value)
}

func (k *Keeper) SetCollateralAuction(ctx sdk.Context, auction auctiontypes.CollateralAuction) {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionKey(auction.Id)
		value = k.cdc.MustMarshal(&auction)
	)
	store.Set(key, value)
}

func (k *Keeper) DeleteCollateralAuction(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionKey(id)
	)
	store.Delete(key)
}

func (k *Keeper) GetCollateralAuction(ctx sdk.Context, id uint64) (auction auctiontypes.CollateralAuction, found bool) {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return auction, false
	}

	k.cdc.MustUnmarshal(value, &auction)
	return auction, true
}

func (k *Keeper) GetCollateralAuctions(ctx sdk.Context) (auctions []auctiontypes.CollateralAuction) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, auctiontypes.CollateralAuctionKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var auction auctiontypes.CollateralAuction
		k.cdc.MustUnmarshal(iter.Value(), &auction)
		auctions = append(auctions, auction)
	}

	return auctions
}

func (k Keeper) PlaceBid(ctx sdk.Context, auctionId uint64, bidder sdk.AccAddress, bid sdk.Coin) error {
	auction, found := k.GetCollateralAuction(ctx, auctionId)
	if !found {
		return auctiontypes.ErrorInvalidAuctionId
	}
	if bid.Denom != auction.Debt.Denom {
		return auctiontypes.ErrorInvalidBiddingDenom
	}
	if bid.Amount.LT(auction.MinBid.Amount) {
		return auctiontypes.ErrorLowBidAmount
	}
	if bid.Amount.GT(auction.MaxBid.Amount) {
		return auctiontypes.ErrorMaxBidAmount
	}
	if bid.Amount.LT(auction.Bid.Amount.Add(sdk.NewInt(1))) {
		return auctiontypes.ErrorBidAlreadyExists
	}
	err := k.SendCoinsFromAccountToModule(ctx, bidder, liquidationtypes.ModuleName, sdk.NewCoins(bid))
	if err != nil {
		return err
	}
	err = k.bank.SendCoinsFromModuleToAccount(ctx, liquidationtypes.ModuleName, auction.Bidder, sdk.NewCoins(auction.Bid))
	if err != nil {
		return err
	}
	auction.Bidder = bidder
	auction.Bid = bid
	k.SetCollateralAuction(ctx, auction)
	return nil
}
