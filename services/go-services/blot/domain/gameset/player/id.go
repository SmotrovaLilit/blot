package player

import "github.com/google/uuid"

type ErrInvalidID struct {
	ID string
}

func (e ErrInvalidID) Error() string {
	return "invalid player id: " + e.ID
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

func (i ID) Equal(p2 ID) bool {
	return i.value == p2.value
}

func NewID(stringID string) (ID, error) {
	id, err := uuid.Parse(stringID)
	if err != nil {
		return ID{}, ErrInvalidID{stringID} // TODO log original error
	}
	return ID{id}, nil
}

func MustNewID(s string) ID {
	id, err := NewID(s)
	if err != nil {
		panic(err)
	}
	return id
}
