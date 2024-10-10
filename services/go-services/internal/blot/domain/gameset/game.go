package gameset

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/team"
	"blot/internal/blot/domain/user"
	"errors"
)

var ErrSameTeam = errors.New("same team")
var ErrGameNotPlaying = errors.New("game is not playing")

type Game struct {
	id        GameID
	gameSetID ID

	team1 Team
	team2 Team

	sittingOrder SittingOrder

	status GameStatus
	round  Round // TODO make optional
	bet    Bet   // TODO make optional
}

func CreateNewGame(
	id GameID,
	gameSetID ID,
	team1 Team,
	team2 Team,
) (*Game, error) {
	if gameSetID.IsZero() || id.IsZero() || team1.IsZero() || team2.IsZero() {
		panic("invalid arguments, create objects using constructors")
	}
	if team1.Equal(team2) {
		return nil, ErrSameTeam
	}
	// TODO check if players unique
	return &Game{
		gameSetID:    gameSetID,
		id:           id,
		team1:        team1,
		team2:        team2,
		status:       GameStatusBetting,
		sittingOrder: NewPlayersSittingOrder(team1.FirstPlayer(), team2.FirstPlayer(), team1.SecondPlayer(), team2.SecondPlayer()),
	}, nil
}

func (g Game) SetBet(bet Bet) {
	if g.status != GameStatusBetting {
		panic("invalid game status")
	}
	//g.round = NewRound(NewRoundNumber(1), []card.Card{})
	g.status = GameStatusPlaying
	g.bet = bet
}

func (g Game) GetTrump() card.Suit {
	if g.status == GameStatusBetting { // TODO remove this check when bet is made optional
		panic("trump is not decided yet")
	}
	return g.bet.suit
}

func (g Game) PlayCard(user user.ID, card card.Card) error {
	if !g.status.CanPlayCard() {
		return ErrGameNotPlaying
	}
	if g.team1.HasPlayer(user) || g.team2.HasPlayer(user) {
		panic("invalid player")
	}
	player := g.findPlayer(user)
	//player.PlayCard(card) moved into round
	g.round.PlayCard(player, card)
	if g.round.Finished() {
		//winner := g.round.Winner(g.GetTrump())
		//winner.AddDiscardCards(g.round.Cards())
	}
	if g.round.IsLastRound() {
		g.finishGame()
		return nil
	}
	return nil
	//g.round = g.round.NextRound()
}

func (g Game) finishGame() {
	if g.status != GameStatusPlaying {
		panic("invalid game status")
	}
	if !g.round.IsLastRound() {
		panic("game is not finished yet")
	}

	g.status = GameStatusFinished
}

func (g Game) Opponent(team Team) team.ID {
	if team.ID() == g.team1.ID() {
		return g.team2.ID()
	}
	if team.ID() == g.team2.ID() {
		return g.team1.ID()
	}
	panic("invalid teamID ")
}

func (g Game) findPlayer(u user.ID) Player {
	if g.team1.HasPlayer(u) {
		return g.team1.findPlayer(u)
	}
	if g.team2.HasPlayer(u) {
		return g.team2.findPlayer(u)
	}
	panic("player not found")
}

func (g Game) StartNewGame(id GameID) (*Game, error) {
	if !g.status.IsFinished() {
		return nil, ErrGameNotFinished
	}
	return CreateNewGame(
		id,
		g.gameSetID,
		NewTeam(g.team1.ID(), g.team1.FirstPlayer().ID(), g.team1.SecondPlayer().ID()),
		NewTeam(g.team2.ID(), g.team2.FirstPlayer().ID(), g.team2.SecondPlayer().ID()),
	)
}
