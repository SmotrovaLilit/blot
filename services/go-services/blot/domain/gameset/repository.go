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

type Repository interface {
	Create(ctx context.Context, gameSet *GameSet) error
	Get(ctx context.Context, id ID) (*GameSet, error)

	// https://threedots.tech/post/database-transactions-in-go/
	UpdateByID(ctx context.Context, setID ID, updateFn func(set *GameSet) (bool, error)) error
}
