package team

import (
	"errors"

	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/player"
	"blot/internal/blot/domain/user"
)

var ErrSamePlayer = errors.New("same player")

type Team struct {
	id      ID
	players [2]player.ID
}

func NewTeam(id ID, p1, p2 player.ID) (Team, error) {
	if id.IsZero() || p1.IsZero() || p2.IsZero() {
		panic("empty input objects, use constructor to create object")
	}
	if p1.Equal(p2) {
		return Team{}, ErrSamePlayer
	}
	return Team{id: id, players: [2]player.ID{p1, p2}}, nil
}

func (t Team) Players() string {
	return t.players[0].String() + " " + t.players[1].String()
}

func (t Team) CardsScore(trump card.Suit) card.Score {
	panic("not implemented")
	// res := card.NewScore(0)
	// for _, p := range t.players {
	//	res = res.Add(p.CalculateScore(trump))
	//}
	// return res
}

// Func (t Team) TeamScore(bet Bet) int {
//	cardsScore := t.CardsScore(bet.suit)
//	if bet.IsFromTeam(t) {
//		if bet.Passed(cardsScore) {
//			return bet.amount + cardsScore.ConvertToTeamScore()
//		}
//		return -bet.amount // TODO ask
//	}
//	return cardsScore.ConvertToTeamScore()
// }.

func (t Team) HasPlayer(u user.ID) bool {
	panic("not implemented")
	// for _, p := range t.players {
	//	if p.userID.Equal(u) {
	//		return true
	//	}
	//}
	// return false
}

// Func (t Team) findPlayer(u user.ID) Player {
//	panic("not implemented")
//	//for _, p := range t.players {
//	//	if p.userID.Equal(u) {
//	//		//return p
//	//	}
//	//}
//	//panic("player not found")
//
// }.

func (t Team) FirstPlayer() player.ID {
	return t.players[0]
}

func (t Team) SecondPlayer() player.ID {
	return t.players[1]
}

func (t Team) Equal(team2 Team) bool {
	return t.players[0].Equal(team2.players[0]) && t.players[1].Equal(team2.players[1])
}

func (t Team) IsZero() bool {
	emptyTeam := Team{}
	return t == emptyTeam
}

func (t Team) ID() ID {
	return t.id
}

func UnmarshalFromDatabase(id ID, p1 player.ID, p2 player.ID) Team {
	return Team{id: id, players: [2]player.ID{p1, p2}}
}
