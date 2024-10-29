package gameset

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/player"
	"github.com/stretchr/testify/require"
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
			expectedError:      game.ErrCardNotFound{PlayerID: "4eb00c05-7f64-47b0-81bf-d0977bff0a04", Card: "ace of Diamonds"}, // TODO fix, it is flaky
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

func prepareGameSetToPlayCard(t *testing.T) *GameSet {
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
	set.MustStartGame(
		game.MustNewID("937cc314-7cf3-4918-8c16-f1699eee89d9"),
		firstPlayerID,
	)
	require.Equal(t, StatusPlaying, set.Status())
	return set
}
