package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/liquidation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
