package command

import (
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/common/decorator"

	"blot/internal/blot/domain/gameset"
)

type StartGame struct {
	SetID    string
	GameID   string
	PlayerID string
}

func (s StartGame) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("set_id", s.SetID),
		slog.String("game_id", s.GameID),
		slog.String("player_id", s.PlayerID),
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
	id, err := gameset.NewID(cmd.SetID)
	if err != nil {
		return err
	}
	gameID, err := game.NewID(cmd.GameID)
	if err != nil {
		return err
	}
	playerID, err := player.NewID(cmd.PlayerID)
	if err != nil {
		return err
	}
	return h.gameSetRepository.UpdateByID(
		ctx,
		id,
		func(set *gameset.GameSet) (bool, error) {
			err := set.StartGame(gameID, playerID)
			if err != nil {
				return false, err
			}
			return true, nil
		},
	)
}
