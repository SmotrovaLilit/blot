package playcard

import (
	"blot/internal/blot/adapters"
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/blot/tests"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayCard_Handle(t *testing.T) {
	testCases := []struct {
		Name               string
		ShouldFail         bool
		ExpectedErrorText  string
		ExpectError        error
		PrepareCommandArgs func(g *gameset.GameSet) PlayCard
		PrepareGameSet     func(repo *adapters.GameSetMemoryRepository) *gameset.GameSet
	}{
		{
			Name: "Should play card",
			PrepareCommandArgs: func(g *gameset.GameSet) PlayCard {
				lastGame := g.LastGame()
				return PlayCard{
					SetID:    g.ID(),
					PlayerID: g.OwnerID(),
					Card:     lastGame.MustPlayerState(g.OwnerID()).HandCards()[0], // TODO how to deal predictable card in deck
				}
			},
			PrepareGameSet: func(repo *adapters.GameSetMemoryRepository) *gameset.GameSet {
				gameSet := tests.PrepareGameSetToPlayCard(t)

				err := repo.Create(context.Background(), gameSet)
				require.NoError(t, err)
				return gameSet
			},
			ShouldFail: false,
		},
		{
			Name: "Should fail when game set is not found",
			PrepareCommandArgs: func(_ *gameset.GameSet) PlayCard {
				return PlayCard{
					SetID:    gameset.MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2"),
					PlayerID: player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04"),
					Card:     card.NewCard(card.RankKing, card.SuitDiamonds),
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
			Name: "Should fail when domain error occurs",
			PrepareCommandArgs: func(g *gameset.GameSet) PlayCard {
				return PlayCard{
					SetID:    g.ID(),
					PlayerID: g.OwnerID(),
					Card:     card.NewCard(card.RankKing, card.SuitDiamonds),
				}
			},
			PrepareGameSet: func(repo *adapters.GameSetMemoryRepository) *gameset.GameSet {
				// Create game that can not be played yet
				gameSet := gameset.NewGameSet(
					gameset.MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2"),
					player.New(player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04"), player.MustNewName("John")),
				)
				err := repo.Create(context.Background(), gameSet)
				require.NoError(t, err)
				return gameSet
			},
			ShouldFail:        true,
			ExpectedErrorText: "game set is not ready to play card",
			ExpectError:       gameset.ErrGameSetNotReadyToPlayCard{Status: gameset.StatusWaitedForPlayers.String()},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			repo := adapters.NewGameSetMemoryRepository()
			h := NewHandler(repo)
			set := tt.PrepareGameSet(repo)
			err := h.Handle(context.Background(), tt.PrepareCommandArgs(set))
			if !tt.ShouldFail {
				require.NoError(t, err)
				_, err := repo.Get(context.Background(), set.ID())
				require.NoError(t, err)
				// TODO check that round is played
			} else {
				require.ErrorContains(t, err, tt.ExpectedErrorText)
				require.ErrorIs(t, err, tt.ExpectError)
			}
		})
	}
}
