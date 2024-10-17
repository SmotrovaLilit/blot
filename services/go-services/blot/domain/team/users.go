package team

import "blot/internal/blot/domain/user"

type Users struct {
	ids [2]user.ID
}

func (u Users) First() user.ID {
	return u.ids[0]
}

func (u Users) Second() user.ID {
	return u.ids[1]
}

func NewUsers(f, s user.ID) Users {
	return Users{[2]user.ID{f, s}}
}
