package query

import (
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/common/decorator"
	"context"
)

type GameSetForPlayer struct {
	GameSetID gameset.ID
	PlayerID  player.ID
}

type GameSetForPlayerQueryHandler decorator.QueryHandler[GameSetForPlayer, gameset.GameSet]

type GameSetForPlayerReadModel interface {
	Get(ctx context.Context, gameSetID gameset.ID) (*gameset.GameSet, error)
}

type gameSetForPlayerQueryHandler struct {
	readModel GameSetForPlayerReadModel
}

func NewGameSetForPlayerQueryHandler(readModel GameSetForPlayerReadModel) GameSetForPlayerQueryHandler {
	if readModel == nil {
		panic("nil readModel")
	}
	return &gameSetForPlayerQueryHandler{readModel: readModel}
}

func (h *gameSetForPlayerQueryHandler) Handle(
	ctx context.Context,
	q GameSetForPlayer,
) (gameset.GameSet, error) {
	s, err := h.readModel.Get(ctx, q.GameSetID)
	if err != nil {
		return gameset.GameSet{}, err
	}

	//if s.Player() != q.PlayerID { // TODO fix it
	//	return gameset.GameSet{}, gameset.NotFoundError{ID: q.GameSetID}
	//}

	return *s, nil
}
