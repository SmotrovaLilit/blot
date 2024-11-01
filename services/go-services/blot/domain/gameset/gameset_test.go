package gameset

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/bet"
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/player"
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"testing"
)

func TestGameSet_PlayCard(t *testing.T) {
	type args struct {
		playerID player.ID
		card     card.Card
	}
	tests := []struct {
		name                  string
		shouldFail            bool
		expectedErrorSting    string
		expectedError         error
		prepareGameSetAndArgs func() (*GameSet, args)
	}{
		{
			name:       "should fail when game in incorrect status",
			shouldFail: true,
			prepareGameSetAndArgs: func() (*GameSet, args) {
				return NewGameSet(
						MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2"),
						player.New(
							player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04"),
							player.MustNewName("John"),
						),
					), args{
						playerID: player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04"),
						card:     card.NewCard(card.RankAce, card.SuitSpades),
					}
			},
			expectedErrorSting: "game set is not ready to play card",
			expectedError:      ErrGameSetNotReadyToPlayCard{Status: StatusWaitedForPlayers.String()},
		},
		{
			name:       "should fail when player not found",
			shouldFail: true,
			prepareGameSetAndArgs: func() (*GameSet, args) {
				set := prepareGameSetToPlayCard(t)
				return set, args{
					playerID: player.MustNewID("4eb00c05-7f88-47b0-8188-d0977bff0a08"),
					card:     card.NewCard(card.RankAce, card.SuitDiamonds),
				}
			},
			expectedErrorSting: "player not found",
			expectedError:      game.ErrPlayerNotFound{ID: player.MustNewID("4eb00c05-7f88-47b0-8188-d0977bff0a08")},
		},
		{
			name:       "should fail when card not found",
			shouldFail: true,
			prepareGameSetAndArgs: func() (*GameSet, args) {
				set := prepareGameSetToPlayCard(t)
				return set, args{
					playerID: player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04"),
					card:     card.NewCard(card.RankAce, card.SuitDiamonds),
				}
			},
			expectedErrorSting: "card not found",
			expectedError:      game.ErrCardNotFound{PlayerID: "4eb00c05-7f64-47b0-81bf-d0977bff0a04", Card: "ace of Diamonds"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set, arg := tt.prepareGameSetAndArgs()
			err := set.PlayCard(arg.playerID, arg.card)
			if tt.shouldFail {
				require.ErrorContains(t, err, tt.expectedErrorSting)
				require.Equal(t, tt.expectedError, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestGameSet_SetBet(t *testing.T) {
	testCases := []struct {
		Name          string
		ShouldFail    bool
		ExpectedError error
		GameSet       *GameSet
	}{
		{
			Name:       "should set bet",
			GameSet:    prepareGameSetToSetBet(t),
			ShouldFail: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			err := tt.GameSet.SetBet(tt.GameSet.ownerID, card.SuitDiamonds, bet.MustNewAmount(8))
			if tt.ShouldFail {
				require.Error(t, err)
				require.Equal(t, tt.ExpectedError, err)
				return
			}
			require.NoError(t, err)
			lastGame := tt.GameSet.LastGame()
			require.Equal(t, card.SuitDiamonds, lastGame.Bet().Trump())
			require.Equal(t, bet.MustNewAmount(8), lastGame.Bet().Amount())
		})
	}
}

func prepareGameSetToStartGame(t *testing.T) *GameSet {
	t.Helper()
	firstPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04")
	secondPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a05")
	thirdPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a06")
	fourthPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a07")
	set := NewGameSet(
		MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2"),
		player.New(
			firstPlayerID,
			player.MustNewName("John"),
		),
	)
	set.MustJoin(player.New(
		secondPlayerID,
		player.MustNewName("Jane"),
	))
	set.MustJoin(player.New(
		thirdPlayerID,
		player.MustNewName("Jack"),
	))
	set.MustJoin(player.New(
		fourthPlayerID,
		player.MustNewName("Jill"),
	))
	require.Equal(t, StatusReadyToStart, set.Status())
	return set
}

func prepareGameSetToPlayCard(t *testing.T) *GameSet {
	t.Helper()
	set := prepareGameSetToSetBet(t)
	set.MustSetBet(set.ownerID, card.SuitSpades, bet.MustNewAmount(8))
	require.Equal(t, StatusPlaying, set.Status())
	return set
}

func prepareGameSetToSetBet(t *testing.T) *GameSet {
	t.Helper()
	set := prepareGameSetToStartGame(t)
	set.MustStartGame(
		game.MustNewID("937cc314-7cf3-4918-8c16-f1699eee89d9"),
		set.ownerID,
		rand.New(rand.NewPCG(0, 0)),
	)
	require.Equal(t, StatusPlaying, set.Status())
	g := set.LastGame()
	require.Equal(t, game.StatusBetting, g.Status())
	return set
}
