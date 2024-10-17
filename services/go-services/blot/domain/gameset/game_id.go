package gameset

import "github.com/google/uuid"

type GameID struct {
	value uuid.UUID
}

func (i GameID) IsZero() bool {
	return i.value == uuid.Nil
}

func (i GameID) String() string {
	return i.value.String()
}

func NewGameID(stringID string) GameID {
	id, err := uuid.Parse(stringID)
	if err != nil {
		panic(err)
	}
	return GameID{id}
}
