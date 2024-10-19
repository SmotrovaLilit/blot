package player

import "fmt"

type Player struct {
	id   ID
	name Name
}

func (p Player) IsZero() bool {
	return p == Player{}
}

func (p Player) ID() ID {
	return p.id
}

func (p Player) Name() Name {
	return p.name
}

func Create(idRaw, nameRaw string) (Player, error) {
	id, err := NewID(idRaw)
	if err != nil {
		return Player{}, err
	}
	name, err := NewName(nameRaw)
	if err != nil {
		return Player{}, err
	}
	return Player{
		id:   id,
		name: name,
	}, nil
}

func New(id ID, name Name) (Player, error) {
	if id.IsZero() || name.IsZero() {
		return Player{}, fmt.Errorf("empty input objects, use constructor to create object")
	}
	return Player{
		id:   id,
		name: name,
	}, nil
}