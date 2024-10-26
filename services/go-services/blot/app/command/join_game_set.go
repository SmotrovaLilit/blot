package command

import (
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset/player"
	"blot/internal/common/decorator"

	"blot/internal/blot/domain/gameset"
)

type JoinGameSet struct {
	Player player.Player
	ID     gameset.ID
}

func (j JoinGameSet) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", j.ID.String()),
		slog.Any("player", j.Player),
	)
}

type joinGameSetHandler struct {
	gameSetRepository gameset.Repository
}

type JoinGameSetHandler decorator.CommandHandler[JoinGameSet]

func NewJoinGameSetHandler(gameSetRepository gameset.Repository) JoinGameSetHandler {
	if gameSetRepository == nil {
		panic("gameSetRepository cannot be nil")
	}
	return decorator.ApplyCommandDecorators(joinGameSetHandler{
		gameSetRepository: gameSetRepository,
	})
}

func (h joinGameSetHandler) Handle(ctx context.Context, cmd JoinGameSet) error {
	return h.gameSetRepository.UpdateByID(ctx, cmd.ID, func(set *gameset.GameSet) (bool, error) {
		err := set.Join(cmd.Player)
		if err != nil {
			return false, err
		}
		return true, nil
	})
}
