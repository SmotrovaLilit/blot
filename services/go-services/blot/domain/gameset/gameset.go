package gameset

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/user"
	"errors"
	"github.com/google/uuid"
)

type GameSet struct {
	id       ID
	lastGame Game
}

var ErrGameNotFinished = errors.New("last game is not finished")

func NewGameSet(id ID, gameID GameID, t1, t2 Team) (*GameSet, error) {
	if id.IsZero() || gameID.IsZero() || t1.IsZero() || t2.IsZero() {
		panic("empty input objects, use constructor to create object")
	}
	game, err := CreateNewGame(gameID, id, t1, t2)
	if err != nil {
		return nil, err
	}
	s := GameSet{
		id:       id,
		lastGame: *game,
	}
	return &s, nil
}

func (s GameSet) ID() ID {
	return s.id
}

func (s GameSet) StartNewGame(gameID GameID) error {
	n, err := s.lastGame.StartNewGame(gameID)
	if err != nil {
		return err
	}
	s.lastGame = *n
	return nil
}

func (s GameSet) PlayCard(id user.ID, card card.Card) error {
	return s.lastGame.PlayCard(id, card)
}

type ID struct {
	value uuid.UUID
}

func (i ID) String() string {
	return i.value.String()
}

func (i ID) IsZero() bool {
	return i.value == uuid.Nil
}

func NewID(stringID string) ID {
	id, err := uuid.Parse(stringID)
	if err != nil {
		panic(err)
	}
	return ID{id}
}
