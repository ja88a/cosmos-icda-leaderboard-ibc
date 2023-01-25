package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/cosmonaut/leaderboard/testutil/keeper"
	"github.com/cosmonaut/leaderboard/testutil/nullify"
	"github.com/cosmonaut/leaderboard/x/leaderboard/keeper"
	"github.com/cosmonaut/leaderboard/x/leaderboard/types"
)

func createTestBoard(keeper *keeper.Keeper, ctx sdk.Context) types.Board {
	item := types.Board{}
	keeper.SetBoard(ctx, item)
	return item
}

func TestBoardGet(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	item := createTestBoard(keeper, ctx)
	rst, found := keeper.GetBoard(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestBoardRemove(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	createTestBoard(keeper, ctx)
	keeper.RemoveBoard(ctx)
	_, found := keeper.GetBoard(ctx)
	require.False(t, found)
}
