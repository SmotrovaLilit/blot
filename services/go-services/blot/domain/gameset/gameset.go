package gameset

import (
	"errors"
	"log/slog"
	"math/rand/v2"
	"strconv"

	"blot/internal/blot/domain/gameset/game/bet"

	"blot/internal/blot/domain/card"

	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/blot/domain/gameset/team"
)

type GameSet struct {
	id       ID
	ownerID  player.ID
	players  []player.Player
	lastGame *game.Game
	status   Status
}

func (s *GameSet) LogValue() slog.Value {
	var players []any
	for i, p := range s.players {
		players = append(players, slog.Any(
			"player"+strconv.Itoa(i+1),
			p,
		))
	}

	return slog.GroupValue(
		slog.String("id", s.id.String()),
		slog.String("status", s.status.String()),
		slog.String("owner_id", s.ownerID.String()),
		slog.Group("players", players...),
	)
}

var ErrGameNotFinished = errors.New("last game is not finished")

func NewGameSet(id ID, pl player.Player) *GameSet {
	if id.IsZero() || pl.IsZero() {
		panic("empty input objects, use constructor to create object")
	}
	s := GameSet{
		id:      id,
		ownerID: pl.ID(),
		players: []player.Player{pl},
		status:  StatusWaitedForPlayers,
	}
	return &s
}

// UnmarshalFromDatabase unmarshals GameSet from the database.
//
// It should be used only for unmarshalling from the database!
// You can't use UnmarshalFromDatabase as constructor - It may put domain into the invalid state!
func UnmarshalFromDatabase(
	id ID,
	status Status,
	ownerID player.ID,
	players []player.Player,
	game game.Game,
) GameSet {
	return GameSet{
		id:       id,
		ownerID:  ownerID,
		players:  players,
		status:   status,
		lastGame: &game,
	}
}

func (s *GameSet) ID() ID {
	return s.id
}

func (s *GameSet) StartGame(gameID game.ID, playerID player.ID, randSource rand.Source) error {
	if gameID.IsZero() || playerID.IsZero() {
		panic("empty input objects, use constructor to create objects")
	}
	if !s.PlayerInGameSet(playerID) {
		return ErrPlayerIsNotInGameSet{playerID}
	}
	newStatus, err := s.status.StartGame()
	if err != nil {
		return err
	}

	team1, err := team.NewTeam(team.MustNewID("1"), s.players[0].ID(), s.players[2].ID())
	if err != nil {
		return err
	}
	team2, err := team.NewTeam(team.MustNewID("2"), s.players[1].ID(), s.players[3].ID())
	if err != nil {
		return err
	}
	lastGame, err := game.NewGame(gameID, team1, team2, randSource)
	if err != nil {
		return err
	}
	s.status = newStatus
	s.lastGame = &lastGame
	// TODO deal cards
	return nil
}

//
// Func (s *GameSet) PlayCard(id user.ID, card card.card) error {
//	return s.lastGame.PlayCard(id, card)
// }.

func (s *GameSet) Status() Status {
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
		s.status = StatusReadyToStart
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

type ErrPlayerIsNotInGameSet struct {
	ID player.ID
}

func (e ErrPlayerIsNotInGameSet) Error() string {
	return "player " + e.ID.String() + " is not in game set"
}

type ErrOwnerCanNotLeaveGameSet struct {
	ID player.ID
}

func (e ErrOwnerCanNotLeaveGameSet) Error() string {
	return "owner " + e.ID.String() + " can not leave game set"
}

func (s *GameSet) RemovePlayer(id player.ID) error {
	if s.OwnerID() == id {
		return ErrOwnerCanNotLeaveGameSet{id}
	}
	for i, p := range s.players {
		if p.ID() == id {
			s.players = append(s.players[:i], s.players[i+1:]...)
			return nil
		}
	}
	return ErrPlayerIsNotInGameSet{id}
}

func (s *GameSet) OwnerID() player.ID {
	return s.ownerID
}

func (s *GameSet) PlayerInGameSet(id player.ID) bool {
	for _, p := range s.players {
		if p.ID() == id {
			return true
		}
	}
	return false
}

// TODO make it optional.
func (s *GameSet) LastGame() game.Game {
	if s.lastGame == nil {
		return game.Game{} // TODO make it optional
	}
	return s.lastGame.Clone()
}

func (s *GameSet) Clone() GameSet {
	players := make([]player.Player, len(s.players))
	copy(players, s.players)
	var g game.Game
	if s.lastGame != nil {
		g = s.lastGame.Clone()
	}
	return GameSet{
		id:       s.id,
		ownerID:  s.ownerID,
		players:  players,
		lastGame: &g,
		status:   s.status,
	}
}

func (s *GameSet) MustJoin(p player.Player) {
	err := s.Join(p)
	if err != nil {
		panic(err)
	}
}

func (s *GameSet) MustStartGame(id game.ID, id2 player.ID, randSource rand.Source) {
	err := s.StartGame(id, id2, randSource)
	if err != nil {
		panic(err)
	}
}

func (s *GameSet) PlayCard(id player.ID, card card.Card) error {
	if !s.status.CanPlayCard() {
		return ErrGameSetNotReadyToPlayCard{s.status.String()}
	}

	return s.lastGame.PlayCard(id, card)
}

func (s *GameSet) SetBet(id player.ID, trump card.Suit, amount bet.Amount) error {
	if !s.status.CanSetBet() {
		return ErrGameSetNotReadyToSetBet{s.status.String()}
	}

	return s.lastGame.SetBet(id, trump, amount)
}

func (s *GameSet) MustSetBet(id player.ID, trump card.Suit, amount bet.Amount) {
	err := s.SetBet(id, trump, amount)
	if err != nil {
		panic(err)
	}
}
