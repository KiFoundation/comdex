package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/liquidation/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	_ types.QueryServer = (*queryServer)(nil)
)

type queryServer struct {
	Keeper
}


func NewQueryServiceServer(k Keeper) types.QueryServer {
	return &queryServer{
		Keeper: k,
	}
}

func (q *queryServer) QueryLockedVault(c context.Context, req *types.QueryLockedVaultRequest) (*types.QueryLockedVaultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	vault, found := q.GetLockedVault(ctx, 1)
	if !found {
		return nil, status.Errorf(codes.NotFound, "vault does not exist for id cool %d", 1)
	}

	assetin,_ := q.asset.GetAsset(ctx, 2)
	assetout,_ := q.asset.GetAsset(ctx,1)
	cr, _:= q.vault.CalculateCollaterlizationRatio(ctx,vault.AmountIn, assetin,vault.AmountOut, assetout)
	return &types.QueryLockedVaultResponse{
		LockedVaultInfo: types.LockedVaultInfo{
			LockedVaultId:          vault.LockedVaultID,
			Id:                     vault.ID,
			PairID:                 vault.PairID,
			Owner:                  vault.Owner,
			Collateral:             sdk.NewCoin("ucGOLD", vault.AmountIn),
			Debt:                   sdk.NewCoin("uatom", vault.AmountOut),
			CollateralizationRatio: cr,
			NewOwner:               vault.NewOwner,
		},
	}, nil
}

func (q *queryServer) QueryAllVaults(c context.Context, req *types.QueryAllVaultsRequest) (*types.QueryAllVaultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		items []types.LockedVaultInfo
		ctx   = sdk.UnwrapSDKContext(c)
	)

	pagination, err := query.FilteredPaginate(
		prefix.NewStore(q.Store(ctx), types.VaultKeyPrefix),
		req.Pagination,
		func(_, value []byte, accumulate bool) (bool, error) {
			var item types.LockedVault
			if err := q.cdc.Unmarshal(value, &item); err != nil {
				return false, err
			}

			pair, found := q.GetPair(ctx, item.PairID)
			if !found {
				return false, status.Errorf(codes.NotFound, "pair does not exist for id %d", item.PairID)
			}

			assetIn, found := q.GetAsset(ctx, pair.AssetIn)
			if !found {
				return false, status.Errorf(codes.NotFound, "asset does not exist for id %d", pair.AssetIn)
			}

			assetOut, found := q.GetAsset(ctx, pair.AssetOut)
			if !found {
				return false, status.Errorf(codes.NotFound, "asset does not exist for id %d", pair.AssetOut)
			}

			collateralizationRatio, err := q.CalculateCollaterlizationRatio(ctx, item.AmountIn, assetIn, item.AmountOut, assetOut)
			if err != nil {
				return false, err
			}

			vaultInfo := types.LockedVaultInfo{
				LockedVaultId: 			item.LockedVaultID,
				Id:                     item.ID,
				PairID:                 item.PairID,
				Owner:                  item.Owner,
				Collateral:             sdk.NewCoin(assetIn.Denom, item.AmountIn),
				Debt:                   sdk.NewCoin(assetOut.Denom, item.AmountOut),
				CollateralizationRatio: collateralizationRatio,
				NewOwner: 				item.NewOwner,
			}

			if accumulate {
				items = append(items, vaultInfo)
			}

			return true, nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVaultsResponse{
		VaultsInfo: items,
		Pagination: pagination,
	}, nil
}
