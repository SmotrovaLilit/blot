package adapters

import (
	"context"

	"blot/internal/blot/domain/gameset"
)

type GameSetPostgresRepository struct {
}

func NewGameSetPostgresRepository() *GameSetPostgresRepository {
	return &GameSetPostgresRepository{}
}

func (g GameSetPostgresRepository) Create(ctx context.Context, gameSet *gameset.GameSet) error {
	//TODO implement me
	panic("implement me")
}

func (g GameSetPostgresRepository) Get(ctx context.Context, id gameset.ID) (*gameset.GameSet, error) {
	//TODO implement me
	panic("implement me")
}

func (g GameSetPostgresRepository) UpdateByID(ctx context.Context, setID gameset.ID, updateFn func(set *gameset.GameSet) (bool, error)) error {
	//TODO implement me
	panic("implement me")
}
