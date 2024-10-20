package gameset

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/blot/domain/user"
	"errors"
	"github.com/google/uuid"
	"log/slog"
	"strconv"
)

type GameSet struct {
	id            ID
	firstPlayerID player.ID
	players       []player.Player
	lastGame      Game
	status        GamesetStatus
}

func (s *GameSet) LogValue() slog.Value {
	var players []interface{}
	for i, p := range s.players {
		players = append(players,
			"player"+strconv.Itoa(i+1), // TODO fix it
			p.LogValue(),
		)
	}

	return slog.GroupValue(
		slog.String("id", s.id.String()),
		slog.String("status", s.status.String()),
		slog.String("first_player_id", s.firstPlayerID.String()),
		slog.Group("players", players...),
	)
}

var ErrGameNotFinished = errors.New("last game is not finished")

func NewGameSet(id ID, pl player.Player) (*GameSet, error) {
	if id.IsZero() || pl.IsZero() {
		panic("empty input objects, use constructor to create object")
	}
	s := GameSet{
		id:            id,
		firstPlayerID: pl.ID(),
		players:       []player.Player{pl},
		status:        GamesetStatusWaitedForPlayers,
	}
	return &s, nil
}

// UnmarshalFromDatabase unmarshals GameSet from the database.
//
// It should be used only for unmarshalling from the database!
// You can't use UnmarshalFromDatabase as constructor - It may put domain into the invalid state!
func UnmarshalFromDatabase(id ID, status GamesetStatus, firstPlayer player.ID, players []player.Player) GameSet {
	return GameSet{
		id:            id,
		firstPlayerID: firstPlayer,
		players:       players,
		status:        status,
	}
}

func (s *GameSet) ID() ID {
	return s.id
}

func (s *GameSet) StartNewGame(gameID GameID) error {
	n, err := s.lastGame.StartNewGame(gameID)
	if err != nil {
		return err
	}
	s.lastGame = *n
	return nil
}

func (s *GameSet) PlayCard(id user.ID, card card.Card) error {
	return s.lastGame.PlayCard(id, card)
}

func (s *GameSet) FirstPlayer() player.Player {
	for _, p := range s.Players() {
		if p.ID() == s.firstPlayerID {
			return p
		}
	}
	panic("first player not found")
}

func (s *GameSet) Status() GamesetStatus {
	return s.status
}

func (s *GameSet) Players() []player.Player {
	return s.players
}

func (s *GameSet) Join(p player.Player) error {
	err := s.CanJoin(p)
	if err != nil {
		return err
	}
	s.players = append(s.players, p)
	if s.isFull() {
		s.status = GamesetStatusReadyToStart
	}
	return nil
}

type ErrGameSetNotAllowJoin struct {
	ID ID
}

func (e ErrGameSetNotAllowJoin) Error() string {
	return "game set " + e.ID.String() + " not allow join"
}

type ErrGameSetFull struct {
	ID ID
}

func (e ErrGameSetFull) Error() string {
	return "game set " + e.ID.String() + " is full"
}

type ErrPlayerAlreadyInGameSet struct {
	ID player.ID
}

func (e ErrPlayerAlreadyInGameSet) Error() string {
	return "player " + e.ID.String() + " already in game set"
}

func (s *GameSet) CanJoin(p player.Player) error {
	if !s.Status().CanJoin() {
		return ErrGameSetNotAllowJoin{s.ID()}
	}
	if s.isFull() {
		return ErrGameSetFull{s.ID()}
	}
	if s.playerInGameSet(p) {
		return ErrPlayerAlreadyInGameSet{p.ID()}
	}
	return nil
}

func (s *GameSet) isFull() bool {
	return len(s.Players()) == 4
}

func (s *GameSet) playerInGameSet(p player.Player) bool {
	for _, pl := range s.Players() {
		if pl.ID() == p.ID() {
			return true
		}
	}
	return false
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
