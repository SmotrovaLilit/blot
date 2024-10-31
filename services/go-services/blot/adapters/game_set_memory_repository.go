package adapters

import (
	"context"
	"log/slog"
	"sync"

	"blot/internal/common/logging"

	"go.opentelemetry.io/otel"

	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/player"
)

var tracer = otel.Tracer("repository")

type GameSetMemoryRepository struct {
	data map[string]gameset.GameSet
	mu   sync.RWMutex
}

func NewGameSetMemoryRepository() *GameSetMemoryRepository {
	return &GameSetMemoryRepository{
		data: make(map[string]gameset.GameSet),
	}
}

func (g *GameSetMemoryRepository) Create(ctx context.Context, gameSet *gameset.GameSet) error {
	ctx, span := tracer.Start(ctx, "gamSetRepo.Create")
	defer span.End()
	ctx = logging.AppendCtx(ctx, slog.String("repo_method", "Create"), slog.Any("set", *gameSet))
	slog.DebugContext(ctx, "repo: creating game set")
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.data[gameSet.ID().String()]; ok {
		slog.ErrorContext(ctx, "repo: failed to create game set: already exists")
		return gameset.ErrGameSetAlreadyExists{ID: gameSet.ID()}
	}
	g.data[gameSet.ID().String()] = gameSet.Clone()
	slog.DebugContext(ctx, "repo: game set created")
	return nil
}

func (g *GameSetMemoryRepository) Get(ctx context.Context, id gameset.ID) (gameset.GameSet, error) {
	ctx, span := tracer.Start(ctx, "gamSetRepo.Get")
	defer span.End()
	ctx = logging.AppendCtx(ctx, slog.String("repo_method", "Get"), slog.String("id", id.String()))
	slog.DebugContext(ctx, "repo: getting game setEntry")
	g.mu.RLock()
	defer g.mu.RUnlock()
	if setEntry, ok := g.data[id.String()]; ok {
		existGame := setEntry.Clone()
		slog.DebugContext(
			ctx,
			"repo: game setEntry found",
			slog.Any("set", existGame),
			slog.Any("set_entry", setEntry),
		)
		return existGame, nil
	}
	slog.DebugContext(ctx, "repo: game setEntry not found")
	return gameset.GameSet{}, gameset.NotFoundError{ID: id}
}

func (g *GameSetMemoryRepository) GetByPlayerID(ctx context.Context, playerID player.ID) ([]gameset.GameSet, error) {
	ctx, span := tracer.Start(ctx, "gamSetRepo.GetByPlayerID")
	defer span.End()
	g.mu.RLock()
	ctx = logging.AppendCtx(ctx, slog.String("repo_method", "GetByPlayerID"), slog.String("player_id", playerID.String()))
	defer g.mu.RUnlock()
	var res []gameset.GameSet
	for _, set := range g.data {
		for _, p := range set.Players() {
			if p.ID().String() == playerID.String() {
				res = append(res, set.Clone())
			}
		}
	}
	slog.DebugContext(ctx, "repo: get player's game set succeed", slog.Any("game_sets", res))
	return res, nil
}
func (g *GameSetMemoryRepository) UpdateByID(ctx context.Context, setID gameset.ID, updateFn func(set *gameset.GameSet) (bool, error)) error {
	ctx, span := tracer.Start(ctx, "gamSetRepo.UpdateByID")
	defer span.End()
	ctx = logging.AppendCtx(ctx, slog.String("repo_method", "UpdateByID"), slog.String("id", setID.String()))
	slog.DebugContext(ctx, "repo: updating game set")
	g.mu.Lock()
	defer g.mu.Unlock()
	setEntry, ok := g.data[setID.String()]
	if !ok {
		return gameset.NotFoundError{ID: setID}
	}
	set := setEntry.Clone()
	ctx = logging.AppendCtx(ctx, slog.Any("set", set), slog.Any("set_entry", setEntry))
	slog.DebugContext(ctx, "repo: got game set to update from memory")
	ok, err := updateFn(&set)
	if err != nil {
		slog.ErrorContext(ctx, "repo: updateFn returns error. set not updated", slog.Any("error", err))
		return err
	}
	if !ok { // we don;t need to  update the set
		slog.DebugContext(ctx, "repo: updateFn returns false, set not updated")
		return nil
	}
	g.data[setID.String()] = set
	slog.DebugContext(ctx, "repo: set updated", slog.Any("set_updated", set), slog.Any("set_entry_updated", setEntry))
	return nil
}
