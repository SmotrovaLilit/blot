package player

import "errors"

var ErrEmptyName = errors.New("empty name")

type Name struct {
	value string
}

func (n Name) IsZero() bool {
	return n == Name{}
}

func (n Name) String() string {
	return n.value
}

func NewName(value string) (Name, error) {
	if value == "" {
		return Name{}, ErrEmptyName
	}
	return Name{value: value}, nil
}
