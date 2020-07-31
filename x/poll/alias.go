package poll

import (
	"github.com/fadeev/poll/x/poll/keeper"
	"github.com/fadeev/poll/x/poll/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	// TODO: Fill out module types
)

var (
	NewMsgCreatePoll = types.NewMsgCreatePoll
)

type (
	MsgCreatePoll = types.MsgCreatePoll
)
		
var (
	NewMsgCreateVote = types.NewMsgCreateVote
)

type (
	MsgCreateVote = types.MsgCreateVote
)
		