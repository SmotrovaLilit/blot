package query

import (
	"context"
	"log/slog"

	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/common/decorator"
)

type GameSetsForPlayer struct {
	PlayerID player.ID
}

func (g GameSetsForPlayer) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("player_id", g.PlayerID.String()),
	)
}

type GameSetsForPlayerQueryHandler decorator.QueryHandler[GameSetsForPlayer, []gameset.GameSet]

type GameSetsForPlayerReadModel interface {
	GetByPlayerID(ctx context.Context, playerID player.ID) ([]gameset.GameSet, error)
}

type gameSetsForPlayerQueryHandler struct {
	readModel GameSetsForPlayerReadModel
}

func NewGameSetsForPlayerQueryHandler(readModel GameSetsForPlayerReadModel) GameSetsForPlayerQueryHandler {
	if readModel == nil {
		panic("nil readModel")
	}
	return decorator.ApplyQueryDecorators(&gameSetsForPlayerQueryHandler{readModel: readModel})
}

func (h *gameSetsForPlayerQueryHandler) Handle(
	ctx context.Context,
	q GameSetsForPlayer,
) ([]gameset.GameSet, error) {
	sets, err := h.readModel.GetByPlayerID(ctx, q.PlayerID)
	if err != nil {
		return nil, err
	}

	return sets, nil
}
