package user

import "github.com/google/uuid"

type User struct {
	id   ID
	name string
}

func NewUser(id ID, name string) User {
	return User{id, name}
}

func (u User) ID() ID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

type ID struct {
	value uuid.UUID
}

func (i ID) String() string {
	return i.value.String()
}

func (i ID) Equal(p2 ID) bool {
	return i.value == p2.value
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
