package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/liquidation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{
		Keeper: keeper,
	}
}

func (k msgServer) MsgLock(c context.Context, request *types.MsgLockRequest) (*types.MsgLockResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)
	_, err := sdk.AccAddressFromBech32(request.From)
	if err != nil {
		return nil, err
	}

	vId := request.Id
	v,_ := k.vault.GetVault(ctx, vId)
	var (
		id  = k.GetLockedVaultID(ctx)
		vault = types.LockedVault{
			LockedVaultID: id + 1,
			ID:            v.ID,
			PairID:        v.PairID,
			Owner:         v.Owner,
			AmountIn:      v.AmountIn,
			AmountOut:     v.AmountOut,
			NewOwner:      request.From,
		}
	)
	k.SetLockedVaultID(ctx, id+1)
	k.SetLockedVault(ctx, vault)
	return &types.MsgLockResponse{}, nil
}
