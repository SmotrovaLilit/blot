package player

import "github.com/google/uuid"

type ErrInvalidID struct {
	value string
}

func (e ErrInvalidID) Error() string {
	return "invalid ID: " + e.value
}

type ID struct {
	value uuid.UUID
}

func (i ID) IsZero() bool {
	return i.value == uuid.Nil
}

func (i ID) String() string {
	return i.value.String()
}

func NewID(stringID string) (ID, error) {
	id, err := uuid.Parse(stringID)
	if err != nil {
		return ID{}, ErrInvalidID{stringID} // TODO log original error
	}
	return ID{id}, nil
}
