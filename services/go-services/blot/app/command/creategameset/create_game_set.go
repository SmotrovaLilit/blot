package creategameset

import (
	"context"

	"blot/internal/blot/domain/gameset/player"

	"blot/internal/blot/domain/gameset"
	"blot/internal/common/decorator"
)

type CreateGameSet struct {
	ID         string
	PlayerID   string
	PlayerName string
}

type gameSetRepository interface {
	Create(ctx context.Context, set *gameset.GameSet) error
}

type createGameSetHandler struct {
	gameSetRepository gameSetRepository
}

type Handler decorator.CommandHandler[CreateGameSet]

func NewHandler(repo gameSetRepository) Handler {
	if repo == nil {
		panic("gameSetRepository cannot be nil")
	}

	return decorator.ApplyCommandDecorators(createGameSetHandler{
		gameSetRepository: repo,
	})
}

func (h createGameSetHandler) Handle(ctx context.Context, cmd CreateGameSet) error {
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
	set := gameset.NewGameSet(id, player.New(playerID, playerName))
	err = h.gameSetRepository.Create(ctx, set)
	if err != nil {
		return err
	}
	return nil
}
