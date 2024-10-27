package command

import (
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset/player"
	"blot/internal/common/decorator"

	"blot/internal/blot/domain/gameset"
)

type JoinGameSet struct {
	ID         string
	PlayerID   string
	PlayerName string
}

func (j JoinGameSet) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", j.ID),
		slog.Any("player_id", j.PlayerID),
		slog.Any("player_name", j.PlayerName),
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
	id, err := gameset.NewID(cmd.ID)
	if err != nil {
		return err
	}
	playerID, err := player.NewID(cmd.PlayerID)
	if err != nil {
		return err
	}
	playerName, err := player.NewName(cmd.PlayerName)
	if err != nil {
		return err
	}
	p := player.New(playerID, playerName)
	return h.gameSetRepository.UpdateByID(ctx, id, func(set *gameset.GameSet) (bool, error) {
		err := set.Join(p)
		if err != nil {
			return false, err
		}
		return true, nil
	})
}
