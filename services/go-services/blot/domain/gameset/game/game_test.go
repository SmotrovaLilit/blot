package game

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/bet"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/blot/domain/gameset/team"
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"testing"
)

func TestGame_SetBet(t *testing.T) {
	testCases := []struct {
		Name          string
		ShouldFail    bool
		ExpectedError error
		Game          *Game
		ArgsPlayerID  player.ID
	}{
		{
			Name:       "should set bet",
			Game:       prepareGameToSetBet(t),
			ShouldFail: false,
		},
		{
			Name:       "should fail when status is not betting",
			ShouldFail: true,
			ExpectedError: ErrGameNotReadyToSetBet{
				Status: StatusPlaying.String(),
			},
			Game: prepareGameToPlayCard(t),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			err := tt.Game.SetBet(tt.Game.players[0].ID(), card.SuitDiamonds, bet.MustNewAmount(8))
			if tt.ShouldFail {
				require.Error(t, err)
				require.Equal(t, tt.ExpectedError, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, card.SuitDiamonds, tt.Game.Bet().Trump())
			require.Equal(t, bet.MustNewAmount(8), tt.Game.Bet().Amount())
		})
	}
	t.Run("should fail when player not found", func(t *testing.T) {
		g := prepareGameToSetBet(t)
		err := g.SetBet(player.MustNewID("4eb00c05-7f88-47b0-8188-d0977bff0a08"), card.SuitDiamonds, bet.MustNewAmount(8))
		require.Error(t, err)
		require.Equal(t, ErrPlayerNotFound{ID: player.MustNewID("4eb00c05-7f88-47b0-8188-d0977bff0a08")}, err)
	})
}

func prepareGameToPlayCard(t *testing.T) *Game {
	g := prepareGameToSetBet(t)
	err := g.SetBet(g.players[0].ID(), card.SuitDiamonds, bet.MustNewAmount(8))
	require.NoError(t, err)
	return g
}

func prepareGameToSetBet(t *testing.T) *Game {
	team1, err := team.NewTeam(team.NewID("1"), player.MustNewID("5667156e-b34e-422e-ab2c-415811b3fbb6"), player.MustNewID("2667156e-b34e-422e-ab2c-415811b3fbb2"))
	require.NoError(t, err)
	team2, err := team.NewTeam(team.NewID("2"), player.MustNewID("3667156e-b34e-422e-ab2c-415811b3fbb3"), player.MustNewID("4667156e-b34e-422e-ab2c-415811b3fbb4"))
	require.NoError(t, err)
	game, err := NewGame(
		MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2"),
		team1,
		team2,
		rand.New(rand.NewPCG(0, 0)),
	)
	require.NoError(t, err)
	return &game
}
