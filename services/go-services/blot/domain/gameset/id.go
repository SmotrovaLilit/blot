package gameset

import (
	"fmt"

	"github.com/google/uuid"
)

type ErrInvalidID struct {
	ID string
}

func (e ErrInvalidID) Error() string {
	return "invalid game set id " + e.ID
}

type ID struct {
	value uuid.UUID
}

func (i ID) String() string {
	return i.value.String()
}

func (i ID) IsZero() bool {
	return i.value == uuid.Nil
}

func NewID(stringID string) (ID, error) {
	id, err := uuid.Parse(stringID)
	if err != nil {
		return ID{}, fmt.Errorf("%w: %v", ErrInvalidID{ID: stringID}, err)
	}
	return ID{id}, nil
}
