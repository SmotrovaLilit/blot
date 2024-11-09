package gameset_test

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/game/bet"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/blot/tests"
	"github.com/stretchr/testify/require"
	"testing"

	. "blot/internal/blot/domain/gameset"
)

func TestGameSet_PlayCard(t *testing.T) {
	type args struct {
		playerID player.ID
		card     card.Card
	}
	testCases := []struct {
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
			name:       "should fail when card not found",
			shouldFail: true,
			prepareGameSetAndArgs: func() (*GameSet, args) {
				set := tests.PrepareGameSetToPlayCard(t)
				return set, args{
					playerID: player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04"),
					card:     card.NewCard(card.RankAce, card.SuitDiamonds),
				}
			},
			expectedErrorSting: "card not found",
			expectedError:      game.ErrCardNotFound{PlayerID: "4eb00c05-7f64-47b0-81bf-d0977bff0a04", Card: "ace of Diamonds"},
		},
	}
	for _, tt := range testCases {
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
			GameSet:    tests.PrepareGameSetToSetBet(t),
			ShouldFail: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			err := tt.GameSet.SetBet(tt.GameSet.OwnerID(), card.SuitDiamonds, bet.MustNewAmount(8))
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
