package command

import (
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/common/decorator"
)

type CreateGameSet struct {
	FirstPlayer player.Player
	ID          gameset.ID
	FirstGameID game.ID
}

func (c CreateGameSet) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", c.ID.String()),
		slog.String("first_player_id", c.FirstGameID.String()),
		slog.Any("first_player", c.FirstPlayer),
	)
}

type createGameSetHandler struct {
	gameSetRepository gameset.Repository
}

type CreateGameSetHandler decorator.CommandHandler[CreateGameSet]

func NewCreateGameSetHandler(gameSetRepository gameset.Repository) CreateGameSetHandler {
	if gameSetRepository == nil {
		panic("gameSetRepository cannot be nil")
	}

	return decorator.ApplyCommandDecorators(createGameSetHandler{
		gameSetRepository: gameSetRepository,
	})
}

func (h createGameSetHandler) Handle(ctx context.Context, cmd CreateGameSet) error {
	set, err := gameset.NewGameSet(
		cmd.ID,
		cmd.FirstPlayer,
	)
	if err != nil {
		return err
	}
	err = h.gameSetRepository.Create(ctx, set)
	if err != nil {
		return err
	}
	return nil
}
