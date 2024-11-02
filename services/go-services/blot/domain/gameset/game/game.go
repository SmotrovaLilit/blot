package game

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"blot/internal/blot/domain/gameset/game/bet"

	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/deck"

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
	rounds []Round
	// rounds  Round // TODO make optional
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
			NewPlayerState(team2.FirstPlayer(), cards[1]),
			NewPlayerState(team1.SecondPlayer(), cards[2]),
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
	rounds := make([]Round, len(g.rounds))
	for i, r := range g.rounds {
		rounds[i] = r.Clone()
	}
	return Game{
		id:      g.id,
		team1:   g.team1,
		team2:   g.team2,
		players: playerStates,
		status:  g.status,
		bet:     g.bet,
		rounds:  rounds,
	}
}

func (g *Game) FirstPlayerState() PlayerState {
	return *g.players[0]
}

func (g *Game) SecondPlayerState() PlayerState {
	return *g.players[1]
}

func (g *Game) ThirdPlayerState() PlayerState {
	return *g.players[2]
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
	err := g.status.CanPlayCard()
	if err != nil {
		return err
	}
	if !g.players[g.currentTurn()].ID().Equal(id) {
		return ErrNotPlayerTurn{
			PlayerID:            id.String(),
			CurrentTurnIndex:    g.currentTurn(),
			CurrentTurnPlayerID: g.players[g.currentTurn()].ID().String(),
			Players:             g.players,
		}
	}
	if len(g.rounds) == 0 {
		panic("rounds not initialized, SetBet maybe works incorrectly")
	}
	round, err := g.rounds[len(g.rounds)-1].PlayCard(NewPlayerCard(id, card))
	if err != nil {
		return err
	}
	err = g.removeCardForPlayer(id, card)
	if err != nil {
		return err
	}
	g.rounds[len(g.rounds)-1] = round
	if round.Finished() {
		number, err := round.Number().Next()
		if err != nil {
			if errors.Is(err, ErrLastRound) {
				g.status = StatusFinished
				return nil
			}
			panic(fmt.Sprintf("unexpected error: %v", err))
		}
		g.rounds = append(g.rounds, NewRound(number, NewTableCards([]PlayerCard{})))
	}
	return nil
}

func (g *Game) removeCardForPlayer(id player.ID, c card.Card) error {
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
	r := NewRound(RoundNumber1, NewTableCards([]PlayerCard{}))
	g.rounds = []Round{r}
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

func (g *Game) LastRound() (Round, error) {
	if len(g.rounds) == 0 {
		return Round{}, errors.New("no rounds")
	}
	return g.rounds[len(g.rounds)-1], nil
}

func (g *Game) currentTurn() int {
	if len(g.rounds) == 0 {
		return 0
	}
	lastRound := g.rounds[len(g.rounds)-1]
	lastRoundNumber := lastRound.Number().Value()
	return calculateTurn(lastRoundNumber, lastRound.Table().Len())
}

func (g *Game) CurrentTurnPlayerID() player.ID {
	return g.players[g.currentTurn()].ID()
}

func (g *Game) Round(number int) (Round, error) {
	if len(g.rounds) < number {
		return Round{}, errors.New("round not found")
	}
	return g.rounds[number-1], nil
}

func (g *Game) Rounds() []Round {
	return g.rounds
}

func calculateTurn(lastRoundNumber, playedTurnsInLastRound int) int {
	firstPlayerIndexInRound := (lastRoundNumber - 1) % 4
	return (firstPlayerIndexInRound + playedTurnsInLastRound) % 4
}

// Func (g Game) SetBet(bet Bet) {
//	if g.status != StatusBetting {
//		panic("invalid game status")
//	}
//	//g.rounds = NewRound(NewRoundNumber(1), []card.card{})
//	g.status = StatusPlaying
//	g.bet = bet
// }.

// Func (g Game) GetTrump() card.Suit {
//	if g.status == StatusBetting { // TODO remove this check when bet is made optional
//		panic("trump is not decided yet")
//	}
//	return g.bet.suit
// }.

// Func (g Game) PlayCard(user user.ID, card card.card) error {
// if !g.status.CanPlayCard() {
//	return ErrGameNotPlaying
//}
// if g.team1.HasPlayer(user) || g.team2.HasPlayer(user) {
//	panic("invalid player")
//}
// player := g.findPlayer(user)
////player.PlayCard(card) moved into rounds
// g.rounds.PlayCard(player, card)
// if g.rounds.Finished() {
//	//winner := g.rounds.Winner(g.GetTrump())
//	//winner.AddDiscardCards(g.rounds.Cards())
//}
// if g.rounds.IsLastRound() {
//	g.finishGame()
//	return nil
//}
// return nil
////g.rounds = g.rounds.NextRound()
// }.

// Func (g Game) finishGame() {
//	if g.status != StatusPlaying {
//		panic("invalid game status")
//	}
//	if !g.rounds.IsLastRound() {
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

type ErrNotPlayerTurn struct {
	PlayerID            string
	CurrentTurnIndex    int
	CurrentTurnPlayerID string
	Players             []*PlayerState
}

func (e ErrNotPlayerTurn) Error() string {
	players := ""
	for _, p := range e.Players {
		players += p.ID().String() + " "
	}
	return fmt.Sprintf("player %s is not in turn, current turn index: %d, current turn player id: %s, players: %s",
		e.PlayerID, e.CurrentTurnIndex, e.CurrentTurnPlayerID, players,
	)
}
