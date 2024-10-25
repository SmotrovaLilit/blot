package command

import (
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/common/decorator"
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset"
)

type StartGame struct {
	SetID    gameset.ID
	GameID   game.ID
	PlayerID player.ID
}

func (s StartGame) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("set_id", s.SetID.String()),
		slog.String("game_id", s.GameID.String()),
		slog.String("player_id", s.PlayerID.String()),
	)
}

type startGameHandler struct {
	gameSetRepository gameset.Repository
}

type StartGameHandler decorator.CommandHandler[StartGame]

func NewStartGameHandler(gameSetRepository gameset.Repository) StartGameHandler {
	if gameSetRepository == nil {
		panic("gameSetRepository cannot be nil")
	}
	return decorator.ApplyCommandDecorators(startGameHandler{
		gameSetRepository: gameSetRepository,
	})
}

func (h startGameHandler) Handle(ctx context.Context, cmd StartGame) error {
	return h.gameSetRepository.UpdateByID(
		ctx,
		cmd.SetID,
		func(set *gameset.GameSet) (bool, error) {
			err := set.StartGame(cmd.GameID, cmd.PlayerID)
			if err != nil {
				return false, err
			}
			return true, nil
		},
	)
}
