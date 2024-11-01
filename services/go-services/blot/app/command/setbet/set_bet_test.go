package setbet

import (
	"blot/internal/blot/adapters"
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/bet"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/blot/tests"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetBetHandler_Handle(t *testing.T) {
	testCases := []struct {
		Name               string
		ShouldFail         bool
		ExpectedErrorText  string
		ExpectError        error
		PrepareCommandArgs func(*gameset.GameSet) SetBet
		PrepareGameSet     func(repo *adapters.GameSetMemoryRepository) *gameset.GameSet
	}{
		{
			Name: "Should set bet",
			PrepareCommandArgs: func(g *gameset.GameSet) SetBet {
				return SetBet{
					SetID:     g.ID(),
					PlayerID:  g.OwnerID(),
					BetTrump:  card.SuitSpades,
					BetAmount: bet.MustNewAmount(8),
				}
			},
			PrepareGameSet: func(repo *adapters.GameSetMemoryRepository) *gameset.GameSet {
				gameSet := tests.PrepareGameSetToSetBet(t)
				err := repo.Create(context.Background(), gameSet)
				require.NoError(t, err)
				return gameSet
			},
			ShouldFail: false,
		},
		{
			Name: "Should fail when game set is not found",
			PrepareCommandArgs: func(_ *gameset.GameSet) SetBet {
				return SetBet{
					SetID:     gameset.MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2"),
					PlayerID:  player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04"),
					BetTrump:  card.SuitSpades,
					BetAmount: bet.MustNewAmount(8),
				}
			},
			PrepareGameSet: func(_ *adapters.GameSetMemoryRepository) *gameset.GameSet {
				return nil
			},
			ShouldFail:        true,
			ExpectedErrorText: "game set '317c8f91-14ef-4582-aaa0-636b5d2ca0c2' not found",
			ExpectError:       gameset.NotFoundError{ID: gameset.MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2")},
		},
		{
			Name: "Should fail when domain error occurred",
			PrepareCommandArgs: func(g *gameset.GameSet) SetBet {
				return SetBet{
					SetID:     g.ID(),
					PlayerID:  player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04"),
					BetTrump:  card.SuitSpades,
					BetAmount: bet.MustNewAmount(8),
				}
			},
			PrepareGameSet: func(repo *adapters.GameSetMemoryRepository) *gameset.GameSet {
				gameSet := gameset.NewGameSet(
					gameset.MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2"),
					player.New(player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04"), player.MustNewName("John")),
				)
				err := repo.Create(context.Background(), gameSet)
				require.NoError(t, err)
				return gameSet
			},
			ShouldFail:        true,
			ExpectedErrorText: "game set is not ready to set bet",
			ExpectError:       gameset.ErrGameSetNotReadyToSetBet{Status: gameset.StatusReadyToStart.String()},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo := adapters.NewGameSetMemoryRepository()
			gameSet := tc.PrepareGameSet(repo)
			handler := NewHandler(repo)
			args := tc.PrepareCommandArgs(gameSet)
			err := handler.Handle(context.Background(), args)
			if tc.ShouldFail {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.ExpectedErrorText)
				return
			}
			require.NoError(t, err)
			s, err := repo.Get(context.Background(), gameSet.ID())
			require.NoError(t, err)
			lastGame := s.LastGame()
			require.Equal(t, args.BetTrump, lastGame.Bet().Trump())
		})
	}
}
