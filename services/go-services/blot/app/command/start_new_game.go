package command

import (
	"blot/internal/common/decorator"
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset"
)

type StartNewGame struct {
	SetID  gameset.ID
	GameID gameset.GameID
}

func (s StartNewGame) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("set_id", s.SetID.String()),
		slog.String("game_id", s.GameID.String()),
	)
}

type startNewGameHandler struct {
	gameSetRepository gameset.Repository
}

type StartNewGameHandler decorator.CommandHandler[StartNewGame]

func NewStartNewGameHandler(gameSetRepository gameset.Repository) StartNewGameHandler {
	if gameSetRepository == nil {
		panic("gameSetRepository cannot be nil")
	}
	return decorator.ApplyCommandDecorators(startNewGameHandler{
		gameSetRepository: gameSetRepository,
	})
}

func (h startNewGameHandler) Handle(ctx context.Context, cmd StartNewGame) error {
	return h.gameSetRepository.UpdateByID(
		ctx,
		cmd.SetID,
		func(set *gameset.GameSet) (bool, error) {
			err := set.StartNewGame(cmd.GameID)
			if err != nil {
				return false, err
			}
			return true, nil
		},
	)
}
