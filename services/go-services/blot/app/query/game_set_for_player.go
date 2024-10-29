package query

import (
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset"
	"blot/internal/common/decorator"
)

type GameSetForPlayer struct {
	GameSetID string
	PlayerID  string
}

func (g GameSetForPlayer) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("game_set_id", g.GameSetID),
		slog.String("player_id", g.PlayerID),
	)
}

type GameSetForPlayerQueryHandler decorator.QueryHandler[GameSetForPlayer, *gameset.GameSet]

type GameSetForPlayerReadModel interface {
	Get(ctx context.Context, id gameset.ID) (gameset.GameSet, error)
}

type gameSetForPlayerQueryHandler struct {
	readModel GameSetForPlayerReadModel
}

func NewGameSetForPlayerQueryHandler(readModel GameSetForPlayerReadModel) GameSetForPlayerQueryHandler {
	if readModel == nil {
		panic("nil readModel")
	}
	return decorator.ApplyQueryDecorators(&gameSetForPlayerQueryHandler{readModel: readModel})
}

func (h *gameSetForPlayerQueryHandler) Handle(
	ctx context.Context,
	q GameSetForPlayer,
) (*gameset.GameSet, error) {
	id, err := gameset.NewID(q.GameSetID)
	if err != nil {
		return nil, err
	}
	s, err := h.readModel.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// if s.Player() != q.PlayerID { // TODO fix it
	//	return gameset.GameSet{}, gameset.NotFoundError{ID: q.GameSetID}
	//}

	return &s, nil
}
