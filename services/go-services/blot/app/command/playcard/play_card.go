package playcard

import (
	"context"

	"blot/internal/blot/domain/gameset/player"

	"blot/internal/common/decorator"

	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset"
)

type PlayCard struct {
	SetID    gameset.ID
	PlayerID player.ID
	Card     card.Card
}

type playCardHandler struct {
	gameSetRepository gameSetRepository
}

type Handler decorator.CommandHandler[PlayCard]

type gameSetRepository interface {
	UpdateByID(ctx context.Context, id gameset.ID, updateFn func(set *gameset.GameSet) (bool, error)) error
}

func NewHandler(gameSetRepository gameSetRepository) Handler {
	if gameSetRepository == nil {
		panic("gameSetRepository cannot be nil")
	}
	return decorator.ApplyCommandDecorators(playCardHandler{
		gameSetRepository: gameSetRepository,
	})
}

func (h playCardHandler) Handle(ctx context.Context, cmd PlayCard) error {
	return h.gameSetRepository.UpdateByID(
		ctx,
		cmd.SetID,
		func(set *gameset.GameSet) (bool, error) {
			err := set.PlayCard(cmd.PlayerID, cmd.Card)
			if err != nil {
				return false, err
			}
			return true, nil
		},
	)
}
