package command

import (
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/common/decorator"
	"context"

	"blot/internal/blot/domain/gameset"
)

type CreateGameSet struct {
	FirstPlayer player.Player
	ID          gameset.ID
	FirstGameID gameset.GameID
}

type createGameSetHandler struct {
	gameSetRepository gameset.Repository
}

type CreateGameSetHandler decorator.CommandHandler[CreateGameSet]

func NewCreateGameSetHandler(gameSetRepository gameset.Repository) CreateGameSetHandler {
	if gameSetRepository == nil {
		panic("gameSetRepository cannot be nil")
	}
	return createGameSetHandler{
		gameSetRepository: gameSetRepository,
	}
}

func (h createGameSetHandler) Handle(ctx context.Context, cmd CreateGameSet) error {
	set, err := gameset.NewGameSet(
		cmd.ID,
		cmd.FirstPlayer,
	)
	err = h.gameSetRepository.Create(ctx, set)
	if err != nil {
		return err
	}
	return nil
}
