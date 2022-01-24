package keeper

import (
	"fmt"
	"github.com/comdex-official/comdex/x/liquidation/expected"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/comdex-official/comdex/x/liquidation/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc        codec.BinaryCodec
	key        sdk.StoreKey
	bank       expected.BankKeeper
	asset      expected.AssetKeeper
	oracle     expected.OracleKeeper
	vault      expected.VaultKeeper
	paramstore paramtypes.Subspace
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, bank expected.BankKeeper, asset expected.AssetKeeper, oracle expected.OracleKeeper, vault expected.VaultKeeper, paramstore paramtypes.Subspace) Keeper {
	return Keeper{
		cdc:        cdc,
		key:        key,
		bank:       bank,
		asset:      asset,
		oracle:     oracle,
		vault:      vault,
		paramstore: paramstore,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.key)
}
