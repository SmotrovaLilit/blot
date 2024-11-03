package game

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/game/bet"
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
			require.Equal(t, StatusPlaying, tt.Game.Status())
			require.Equal(t, RoundNumber1, tt.Game.rounds[0].number)
			require.Equal(t, tt.Game.team1.ID(), tt.Game.bet.TeamID())
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
	team1, err := team.NewTeam(team.MustNewID("1"), player.MustNewID("5667156e-b34e-422e-ab2c-415811b3fbb6"), player.MustNewID("2667156e-b34e-422e-ab2c-415811b3fbb2"))
	require.NoError(t, err)
	team2, err := team.NewTeam(team.MustNewID("2"), player.MustNewID("3667156e-b34e-422e-ab2c-415811b3fbb3"), player.MustNewID("4667156e-b34e-422e-ab2c-415811b3fbb4"))
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

func TestGame_PlayCard(t *testing.T) {
	playingGame := prepareGameToPlayCard(t)
	bettingGame := prepareGameToSetBet(t)
	testCases := []struct {
		Name          string
		ShouldFail    bool
		ExpectedError error
		Game          Game
		ArgsPlayerID  player.ID
		ArgsCard      card.Card
	}{
		{
			Name:         "should play card",
			Game:         playingGame.Clone(),
			ShouldFail:   false,
			ArgsPlayerID: playingGame.players[0].ID(),
			ArgsCard:     playingGame.players[0].HandCards()[0],
		},
		{
			Name:         "should fail when game is not playing",
			ShouldFail:   true,
			Game:         bettingGame.Clone(),
			ArgsPlayerID: bettingGame.players[0].ID(),
			ArgsCard:     bettingGame.players[0].HandCards()[0],
			ExpectedError: ErrGameNotReadyToPlayCard{
				Status: StatusBetting.String(),
			},
		},
		{
			Name:         "should fail when not player's turn",
			ShouldFail:   true,
			Game:         playingGame.Clone(),
			ArgsPlayerID: playingGame.players[1].ID(),
			ArgsCard:     playingGame.players[1].HandCards()[0],
			ExpectedError: ErrNotPlayerTurn{
				PlayerID:            playingGame.players[1].ID().String(),
				CurrentTurnIndex:    0,
				CurrentTurnPlayerID: playingGame.players[0].ID().String(),
				Players:             playingGame.players,
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			err := tt.Game.PlayCard(tt.ArgsPlayerID, tt.ArgsCard)
			if tt.ShouldFail {
				require.Error(t, err)
				require.Equal(t, tt.ExpectedError, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, StatusPlaying, tt.Game.Status())
			require.Len(t, tt.Game.players[0].HandCards(), 7)
			require.True(t, tt.Game.rounds[0].HasCard(tt.ArgsCard))
		})
	}
}

func Test_calculateTurn(t *testing.T) {
	type args struct {
		lastRoundNumber        int
		playedTurnsInLastRound int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should calculate turn 0",
			args: args{
				lastRoundNumber:        1,
				playedTurnsInLastRound: 0,
			},
			want: 0,
		},
		{
			name: "should calculate turn 1",
			args: args{
				lastRoundNumber:        1,
				playedTurnsInLastRound: 1,
			},
			want: 1,
		},
		{
			name: "should calculate turn 2",
			args: args{
				lastRoundNumber:        1,
				playedTurnsInLastRound: 2,
			},
			want: 2,
		},
		{
			name: "should calculate turn 3",
			args: args{
				lastRoundNumber:        1,
				playedTurnsInLastRound: 3,
			},
			want: 3,
		},
		{
			name: "should calculate turn 1 in second round",
			args: args{
				lastRoundNumber:        2,
				playedTurnsInLastRound: 0,
			},
			want: 1,
		},
		{
			name: "should calculate turn 2 in second round",
			args: args{
				lastRoundNumber:        2,
				playedTurnsInLastRound: 1,
			},
			want: 2,
		},
		{
			name: "should calculate turn 3 in second round",
			args: args{
				lastRoundNumber:        2,
				playedTurnsInLastRound: 2,
			},
			want: 3,
		},
		{
			name: "should calculate turn 0 in second round",
			args: args{
				lastRoundNumber:        2,
				playedTurnsInLastRound: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateTurn(tt.args.lastRoundNumber, tt.args.playedTurnsInLastRound)
			require.Equal(t, tt.want, got)
		})
	}
}
