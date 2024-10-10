package user

import "fmt"

type NotFoundError struct {
	ID ID
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("user '%s' not found", e.ID.String())
}

type Repository interface {
}
