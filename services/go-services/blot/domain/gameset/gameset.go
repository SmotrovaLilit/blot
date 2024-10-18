package gameset

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/blot/domain/user"
	"errors"
	"github.com/google/uuid"
)

type GameSet struct {
	id          ID
	firstPlayer player.Name
	lastGame    Game
	status      GamesetStatus
}

var ErrGameNotFinished = errors.New("last game is not finished")

func NewGameSet(id ID, name player.Name) (*GameSet, error) {
	if id.IsZero() || name.IsZero() {
		panic("empty input objects, use constructor to create object")
	}
	s := GameSet{
		id:          id,
		firstPlayer: name,
		status:      GamesetStatusWaitedForPlayers,
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

func (s GameSet) FirstPlayer() player.Name {
	return s.firstPlayer
}

func (s GameSet) Status() GamesetStatus {
	return s.status
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
