package gameset

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	ID ID
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("game set '%s' not found", e.ID.String())
}

type ErrGameSetAlreadyExists struct {
	ID ID
}

func (e ErrGameSetAlreadyExists) Error() string {
	return fmt.Sprintf("game set '%s' already exists", e.ID.String())
}

type Repository interface {
	Create(ctx context.Context, gameSet *GameSet) error

	// https://threedots.tech/post/database-transactions-in-go/
	UpdateByID(ctx context.Context, setID ID, updateFn func(set *GameSet) (bool, error)) error
}
