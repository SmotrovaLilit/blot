package team

type ID struct {
	value string
}

func NewID(stringID string) (ID, error) {
	if stringID == "" {
		return ID{}, ErrInvalidID{stringID}
	}
	return ID{stringID}, nil
}

func MustNewID(s string) ID {
	id, err := NewID(s)
	if err != nil {
		panic(err)
	}
	return id
}

func (i ID) IsZero() bool {
	return i == ID{}
}

func (i ID) String() string {
	return i.value
}

type ErrInvalidID struct {
	ID string
}

func (e ErrInvalidID) Error() string {
	return "invalid team id: " + e.ID
}
