package gameset

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/team"
	"blot/internal/blot/domain/user"
	"errors"
)

var ErrSamePlayer = errors.New("same player")

type Team struct {
	id      team.ID
	players [2]*Player
}

func (t Team) ID() team.ID {
	return t.id
}

func NewTeam(t team.ID, p1, p2 user.ID) Team {
	return Team{id: t, players: [2]*Player{
		NewPlayer(p1, t),
		NewPlayer(p2, t),
	}}
}

func (t Team) CardsScore(trump card.Suit) card.Score {
	res := card.NewScore(0)
	for _, p := range t.players {
		res = res.Add(p.CalculateScore(trump))
	}
	return res
}

func (t Team) TeamScore(bet Bet) int {
	cardsScore := t.CardsScore(bet.suit)
	if bet.IsFromTeam(t) {
		if bet.Passed(cardsScore) {
			return bet.amount + cardsScore.ConvertToTeamScore()
		}
		return -bet.amount // TODO ask
	}
	return cardsScore.ConvertToTeamScore()
}

func (t Team) HasPlayer(u user.ID) bool {
	for _, p := range t.players {
		if p.userID.Equal(u) {
			return true
		}
	}
	return false
}

func (t Team) findPlayer(u user.ID) Player {
	for _, p := range t.players {
		if p.userID.Equal(u) {
			//return p
		}
	}
	panic("player not found")

}

func (t Team) FirstPlayer() *Player {
	return t.players[0]
}

func (t Team) SecondPlayer() *Player {
	return t.players[1]
}

func (t Team) Equal(team2 Team) bool {
	return t.id == team2.id
}

func (t Team) IsZero() bool {
	emptyTeam := Team{}
	return t == emptyTeam
}
