package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fadeev/poll/x/poll/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateVote(ctx sdk.Context, vote types.Vote) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.VotePrefix + vote.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(vote)
	store.Set(key, value)
}

func listVote(ctx sdk.Context, k Keeper) ([]byte, error) {
  var voteList []types.Vote
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.VotePrefix))
  for ; iterator.Valid(); iterator.Next() {
    var vote types.Vote
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &vote)
    voteList = append(voteList, vote)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, voteList)
  return res, nil
}