package setbet

import (
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset/game/bet"

	"blot/internal/common/decorator"

	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
)

type SetBet struct {
	SetID     gameset.ID
	PlayerID  player.ID
	BetTrump  card.Suit
	BetAmount bet.Amount
}

func (s SetBet) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("set_id", s.SetID.String()),
		slog.String("player_id", s.PlayerID.String()),
		slog.String("bet_trump", s.BetTrump.String()))
}

type Handler decorator.CommandHandler[SetBet]

type GameSetRepository interface {
	UpdateByID(ctx context.Context, setID gameset.ID, updateFn func(set *gameset.GameSet) (bool, error)) error
}

type setBetHandler struct {
	gameSetRepository GameSetRepository
}

func NewHandler(gameSetRepository GameSetRepository) Handler {
	if gameSetRepository == nil {
		panic("gameSetRepository cannot be nil")
	}
	return setBetHandler{
		gameSetRepository: gameSetRepository,
	}
}

func (h setBetHandler) Handle(ctx context.Context, cmd SetBet) error {
	return h.gameSetRepository.UpdateByID(
		ctx,
		cmd.SetID,
		func(set *gameset.GameSet) (bool, error) {
			err := set.SetBet(cmd.PlayerID, cmd.BetTrump, cmd.BetAmount)
			if err != nil {
				return false, err
			}
			return true, nil
		},
	)
}
