package command

import (
	"blot/internal/common/decorator"
	"context"

	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/user"
)

type PlayCard struct {
	SetID    gameset.ID
	PlayerID user.ID
	Card     card.Card
}

type playCardHandler struct {
	gameSetRepository gameset.Repository
}

type PlayCardHandler decorator.CommandHandler[PlayCard]

func NewPlayCardHandler(gameSetRepository gameset.Repository) PlayCardHandler {
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
