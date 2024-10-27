package player

import (
	"log/slog"
)

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

func (p Player) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", p.id.String()),
		slog.String("name", p.name.String()),
	)
}

func UnmarshalFromDatabase(idRaw, nameRaw string) (Player, error) {
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

func New(id ID, name Name) Player {
	if id.IsZero() || name.IsZero() {
		panic("empty input objects, use constructor to create object")
	}
	return Player{
		id:   id,
		name: name,
	}
}
