package game

import (
	"errors"
	"math/rand/v2"

	"blot/internal/blot/domain/deck"
	"blot/internal/blot/domain/gameset/bet"

	"blot/internal/blot/domain/card"

	"blot/internal/blot/domain/gameset/player"

	"blot/internal/blot/domain/gameset/team"
)

var ErrSameTeam = errors.New("same team")

type ErrPlayerNotFound struct {
	ID player.ID
}

func (e ErrPlayerNotFound) Error() string {
	return "player not found: " + e.ID.String()
}

// Var ErrGameNotPlaying = errors.New("game is not playing").

type Game struct {
	id ID

	team1   team.Team
	team2   team.Team
	players []*PlayerState

	// sittingOrder SittingOrder

	status Status
	bet    bet.Bet
	// round  Round // TODO make optional
	// bet    Bet   // TODO make optional
}

func NewGame(
	id ID,
	team1 team.Team,
	team2 team.Team,
	randSource rand.Source,
) (Game, error) {
	if id.IsZero() || team1.IsZero() || team2.IsZero() {
		panic("invalid arguments, create objects using constructors")
	}
	if team1.Equal(team2) {
		return Game{}, ErrSameTeam
	}

	if team1.FirstPlayer().Equal(team2.SecondPlayer()) || team1.FirstPlayer().Equal(team2.FirstPlayer()) {
		panic("same player in different teams")
	}
	cards := deck.NewDeck(randSource).DealCards()
	return Game{
		id:     id,
		team1:  team1,
		team2:  team2,
		status: StatusBetting,
		players: []*PlayerState{
			NewPlayerState(team1.FirstPlayer(), cards[0]),
			NewPlayerState(team2.FirstPlayer(), cards[2]),
			NewPlayerState(team1.SecondPlayer(), cards[1]),
			NewPlayerState(team2.SecondPlayer(), cards[3]),
		},
		//sittingOrder: NewPlayersSittingOrder(team1.FirstPlayer(), team2.FirstPlayer(), team1.SecondPlayer(), team2.SecondPlayer()),
	}, nil
}

func (g *Game) ID() ID {
	return g.id
}

func (g *Game) Status() Status {
	return g.status
}

func (g *Game) FirstTeam() team.Team {
	return g.team1
}

func (g *Game) SecondTeam() team.Team {
	return g.team2
}

func (g *Game) PlayerStates() []PlayerState {
	playerStates := make([]PlayerState, len(g.players))
	for i, p := range g.players {
		playerStates[i] = *p
	}
	return playerStates
}

func (g *Game) IsZero() bool {
	return g.status.IsZero()
}

func (g *Game) Clone() Game {
	playerStates := make([]*PlayerState, len(g.players))
	for i, p := range g.players {
		newP := p.Clone()
		playerStates[i] = &newP
	}
	return Game{
		id:      g.id,
		team1:   g.team1,
		team2:   g.team2,
		players: playerStates,
		status:  g.status,
		bet:     g.bet,
	}
}

func (g *Game) FirstPlayerState() PlayerState {
	return *g.players[0]
}

func (g *Game) MustPlayerState(id player.ID) PlayerState {
	p, err := g.PlayerState(id)
	if err != nil {
		panic(err)
	}
	return p
}

func (g *Game) PlayerState(id player.ID) (PlayerState, error) {
	for _, p := range g.players {
		if p.ID() == id {
			return p.Clone(), nil
		}
	}
	return PlayerState{}, ErrPlayerNotFound{ID: id}
}

func (g *Game) PlayCard(id player.ID, card card.Card) error {
	if id.IsZero() || card.IsZero() {
		panic("invalid arguments, create objects using constructors")
	}
	newStatus, err := g.status.PlayCard()
	if err != nil {
		return err
	}
	err = g.RemoveCardForPlayer(id, card)
	if err != nil {
		return err
	}
	g.status = newStatus
	return nil
}

func (g *Game) RemoveCardForPlayer(id player.ID, c card.Card) error {
	for i, p := range g.players { // TODO refactor to use iterator
		if p.ID() == id {
			return g.players[i].RemoveCard(c)
		}
	}
	return ErrPlayerNotFound{ID: id}
}

func (g *Game) SetBet(id player.ID, trump card.Suit, amount bet.Amount) error {
	if id.IsZero() || trump.IsZero() || amount.IsZero() {
		// TODO think about returning error instead of panic
		panic("invalid arguments, create objects using constructors")
	}
	newStatus, err := g.status.SetBet()
	if err != nil {
		return err
	}
	teamID, err := g.TeamByPlayerID(id)
	if err != nil {
		return err
	}
	g.bet = bet.NewBet(teamID, amount, trump)
	g.status = newStatus
	return nil
}

func (g *Game) TeamByPlayerID(id player.ID) (team.ID, error) {
	if g.team1.HasPlayer(id) {
		return g.team1.ID(), nil
	}
	if g.team2.HasPlayer(id) {
		return g.team2.ID(), nil
	}
	return team.ID{}, ErrPlayerNotFound{ID: id}
}

func (g *Game) Bet() bet.Bet {
	return g.bet
}

// Func (g Game) SetBet(bet Bet) {
//	if g.status != StatusBetting {
//		panic("invalid game status")
//	}
//	//g.round = NewRound(NewRoundNumber(1), []card.Card{})
//	g.status = StatusPlaying
//	g.bet = bet
// }.

// Func (g Game) GetTrump() card.Suit {
//	if g.status == StatusBetting { // TODO remove this check when bet is made optional
//		panic("trump is not decided yet")
//	}
//	return g.bet.suit
// }.

// Func (g Game) PlayCard(user user.ID, card card.Card) error {
// if !g.status.CanPlayCard() {
//	return ErrGameNotPlaying
//}
// if g.team1.HasPlayer(user) || g.team2.HasPlayer(user) {
//	panic("invalid player")
//}
// player := g.findPlayer(user)
////player.PlayCard(card) moved into round
// g.round.PlayCard(player, card)
// if g.round.Finished() {
//	//winner := g.round.Winner(g.GetTrump())
//	//winner.AddDiscardCards(g.round.Cards())
//}
// if g.round.IsLastRound() {
//	g.finishGame()
//	return nil
//}
// return nil
////g.round = g.round.NextRound()
// }.

// Func (g Game) finishGame() {
//	if g.status != StatusPlaying {
//		panic("invalid game status")
//	}
//	if !g.round.IsLastRound() {
//		panic("game is not finished yet")
//	}
//
//	g.status = StatusFinished
// }.

// Func (g Game) Opponent(team Team) team.ID {
//	panic("not implemented")
//	//if team.ID() == g.team1.ID() {
//	//	return g.team2.ID()
//	//}
//	//if team.ID() == g.team2.ID() {
//	//	return g.team1.ID()
//	//}
//	//panic("invalid teamID ")
//}
//
// func (g Game) findPlayer(u user.ID) Player {
//	if g.team1.HasPlayer(u) {
//		return g.team1.findPlayer(u)
//	}
//	if g.team2.HasPlayer(u) {
//		return g.team2.findPlayer(u)
//	}
//	panic("player not found")
// }.

// Func (g Game) StartNewGame(id ID) (*Game, error) {
//	panic("not implemented")
//	//if !g.status.IsFinished() {
//	//	return nil, ErrGameNotFinished
//	//}
//	//return NewGame(
//	//	id,
//	//	g.gameSetID,
//	//	NewTeam(g.team1.ID(), g.team1.FirstPlayer().ID(), g.team1.SecondPlayer().ID()),
//	//	NewTeam(g.team2.ID(), g.team2.FirstPlayer().ID(), g.team2.SecondPlayer().ID()),
//	//)
// }.

func UnmarshalFromDatabase(id ID, status Status, team1 team.Team, team2 team.Team, states []PlayerState) Game {
	playerStates := make([]*PlayerState, len(states))
	for i, s := range states {
		playerStates[i] = &s
	}
	return Game{
		id:      id,
		status:  status,
		team1:   team1,
		team2:   team2,
		players: playerStates,
	}
}
