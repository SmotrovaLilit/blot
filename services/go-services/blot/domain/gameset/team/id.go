package team

type ID struct {
	value string
}

func (i ID) IsZero() bool {
	return i == ID{}
}

func (i ID) String() string {
	return i.value
}

func NewID(stringID string) ID {
	// TODO validate
	return ID{stringID}
}
