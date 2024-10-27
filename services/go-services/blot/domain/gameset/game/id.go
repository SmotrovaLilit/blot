package game

import "github.com/google/uuid"

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
		return ID{}, err
	}
	return ID{id}, nil
}
