package team

import (
	"blot/internal/blot/domain/user"

	"github.com/google/uuid"
)

type Team struct {
	id      ID
	name    Name
	userIds Users
}

func NewTeam(id ID, name Name, users Users) Team {
	return Team{id: id, name: name, userIds: users}
}

func (t Team) Equal(s Team) bool {
	return t.id == s.id
}

func (t Team) ID() ID {
	return t.id
}

func (t Team) FirstPlayerID() user.ID {
	return t.userIds.First()
}

func (t Team) SecondPlayerID() user.ID {
	return t.userIds.Second()
}

type ID struct {
	value uuid.UUID
}

func (i ID) String() string {
	return i.value.String()
}

func (i ID) Equal(s ID) bool {
	return i.value == s.value
}

func (i ID) IsZero() bool {
	return i.value == uuid.UUID{}
}

func NewID(stringID string) ID {
	id, err := uuid.Parse(stringID)
	if err != nil {
		panic(err)
	}
	return ID{id}
}
