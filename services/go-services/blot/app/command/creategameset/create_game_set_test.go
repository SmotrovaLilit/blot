package creategameset

import (
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_createGameSetHandler_Handle(t *testing.T) {
	type args struct {
		ID         string
		PlayerID   string
		PlayerName string
	}
	testCases := []struct {
		Name              string
		ShouldFail        bool
		ExpectedErrorText string
		ExpectError       error
		Args              args
	}{
		{
			Name:       "Should create game set",
			Args:       args{ID: "317c8f91-14ef-4582-aaa0-636b5d2ca0c2", PlayerID: "4eb00c05-7f64-47b0-81bf-d0977bff0a04", PlayerName: "John"},
			ShouldFail: false,
		},
		{
			Name:              "Should fail when game set id is invalid",
			Args:              args{ID: "317c8f91-14ef-4582-aaa0-636b5d2c", PlayerID: "4eb00c05-7f64-47b0-81bf-d0977bff0a04", PlayerName: "John"},
			ShouldFail:        true,
			ExpectedErrorText: "invalid game set id",
			ExpectError:       gameset.ErrInvalidID{ID: "317c8f91-14ef-4582-aaa0-636b5d2c"},
		},
		{
			Name:              "Should fail when player id is invalid",
			Args:              args{ID: "317c8f91-14ef-4582-aaa0-636b5d2ca0c2", PlayerID: "4eb00c05-7f64-47b0-81bf-d097", PlayerName: "John"},
			ShouldFail:        true,
			ExpectedErrorText: "invalid player id",
			ExpectError:       player.ErrInvalidID{ID: "4eb00c05-7f64-47b0-81bf-d097"},
		},
		{
			Name:              "Should fail when player name is invalid",
			Args:              args{ID: "317c8f91-14ef-4582-aaa0-636b5d2ca0c2", PlayerID: "4eb00c05-7f64-47b0-81bf-d0977bff0a04", PlayerName: ""},
			ShouldFail:        true,
			ExpectedErrorText: "empty player name",
			ExpectError:       player.ErrEmptyName,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			repo := nwGameSetRepositoryMock(nil)
			h := NewHandler(repo)
			err := h.Handle(context.Background(), CreateGameSet{
				ID:         tt.Args.ID,
				PlayerID:   tt.Args.PlayerID,
				PlayerName: tt.Args.PlayerName,
			})
			if tt.ShouldFail {
				require.ErrorContains(t, err, tt.ExpectedErrorText)
				require.ErrorIs(t, err, tt.ExpectError)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "317c8f91-14ef-4582-aaa0-636b5d2ca0c2", repo.createdGameSet.ID().String())
			assert.Equal(t, "4eb00c05-7f64-47b0-81bf-d0977bff0a04", repo.createdGameSet.Players()[0].ID().String())
			assert.Equal(t, "John", repo.createdGameSet.Players()[0].Name().String())
			assert.Equal(t, gameset.StatusWaitedForPlayers, repo.createdGameSet.Status())
		})
	}
	t.Run("Should fail when repository fails", func(t *testing.T) {
		t.Parallel()
		errSome := errors.New("some error")
		repo := nwGameSetRepositoryMock(errSome)
		h := NewHandler(repo)
		err := h.Handle(context.Background(), CreateGameSet{
			ID:         "317c8f91-14ef-4582-aaa0-636b5d2ca0c2",
			PlayerID:   "4eb00c05-7f64-47b0-81bf-d0977bff0a04",
			PlayerName: "John",
		})
		require.Error(t, err)
		require.ErrorContains(t, err, errSome.Error())
		require.ErrorIs(t, err, errSome)

	})
}

type gameSetRepositoryMock struct {
	createdGameSet *gameset.GameSet
	returnErr      error
}

func nwGameSetRepositoryMock(returnErr error) *gameSetRepositoryMock {
	return &gameSetRepositoryMock{
		returnErr: returnErr,
	}
}

func (g *gameSetRepositoryMock) Create(ctx context.Context, gameSet *gameset.GameSet) error {
	if g.returnErr != nil {
		return g.returnErr
	}
	g.createdGameSet = gameSet
	return nil
}
