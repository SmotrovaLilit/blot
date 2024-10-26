package command

import (
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/common/decorator"
)

type LeaveGameSet struct {
	ID       gameset.ID
	PlayerID player.ID
}

func (l LeaveGameSet) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", l.ID.String()),
		slog.String("player_id", l.PlayerID.String()),
	)
}

type leaveGameSetHandler struct {
	gameSetRepository gameset.Repository
}

type LeaveGameSetHandler decorator.CommandHandler[LeaveGameSet]

func NewLeaveGameSetHandler(gameSetRepository gameset.Repository) LeaveGameSetHandler {
	if gameSetRepository == nil {
		panic("gameSetRepository cannot be nil")
	}

	return decorator.ApplyCommandDecorators(leaveGameSetHandler{
		gameSetRepository: gameSetRepository,
	})
}

func (h leaveGameSetHandler) Handle(ctx context.Context, cmd LeaveGameSet) error {
	return h.gameSetRepository.UpdateByID(ctx, cmd.ID, func(set *gameset.GameSet) (bool, error) {
		// TODO think about business logic later
		err := set.RemovePlayer(cmd.PlayerID)
		if err != nil {
			return false, err
		}
		return true, nil
	})
}
