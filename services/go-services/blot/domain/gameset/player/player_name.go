package player

type Name struct {
	value string
}

func (n Name) IsZero() bool {
	return n == Name{}
}

func NewName(value string) (Name, error) {
	// TODO validation
	return Name{value: value}, nil
}
