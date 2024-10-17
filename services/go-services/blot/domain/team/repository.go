package team

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	ID ID
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("team '%s' not found", e.ID.String())
}

type Repository interface {
	Get(ctx context.Context, id ID) (*Team, error)
}
