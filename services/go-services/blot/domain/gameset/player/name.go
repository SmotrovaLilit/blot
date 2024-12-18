package player

import "errors"

var ErrEmptyName = errors.New("empty player name")

type Name struct {
	value string
}

func (i Name) IsZero() bool {
	return i == Name{}
}

func (i Name) String() string {
	return i.value
}

func NewName(stringValue string) (Name, error) {
	if stringValue == "" {
		return Name{}, ErrEmptyName
	}
	return Name{stringValue}, nil
}

func MustNewName(s string) Name {
	n, err := NewName(s)
	if err != nil {
		panic(err)
	}
	return n
}
