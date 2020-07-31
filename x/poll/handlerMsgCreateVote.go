package poll

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fadeev/poll/x/poll/types"
)

func handleMsgCreateVote(ctx sdk.Context, k Keeper, msg MsgCreateVote) (*sdk.Result, error) {
	var vote = types.Vote{
		Creator: msg.Creator,
		ID:      msg.ID,
    PollID: msg.PollID,
    Value: msg.Value,
	}
	k.CreateVote(ctx, vote)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
