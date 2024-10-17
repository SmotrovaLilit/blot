package adapters

import (
	"context"
	"sync"

	"blot/internal/blot/domain/gameset"
)

type GameSetMemoryRepository struct {
	data map[gameset.ID]*gameset.GameSet
	mu   sync.RWMutex
}

func NewGameSetMemoryRepository() *GameSetMemoryRepository {
	return &GameSetMemoryRepository{
		data: make(map[gameset.ID]*gameset.GameSet),
	}
}

func (g *GameSetMemoryRepository) Create(ctx context.Context, gameSet *gameset.GameSet) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.data[gameSet.ID()]; ok {
		return gameset.ErrGameSetAlreadyExists{ID: gameSet.ID()}
	}
	g.data[gameSet.ID()] = gameSet
	return nil
}

func (g *GameSetMemoryRepository) Get(ctx context.Context, id gameset.ID) (*gameset.GameSet, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	if set, ok := g.data[id]; ok {
		return set, nil
	}
	return nil, gameset.NotFoundError{ID: id}
}

func (g *GameSetMemoryRepository) UpdateByID(ctx context.Context, setID gameset.ID, updateFn func(set *gameset.GameSet) (bool, error)) error {
	//TODO implement me
	panic("implement me")
}
