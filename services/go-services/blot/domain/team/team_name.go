package team

type Name struct {
	value string
}

func NewName(name string) Name {
	return Name{name}
}
