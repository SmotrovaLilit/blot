package bet

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/team"
)

type Bet struct {
	teamID team.ID
	amount int
	suit   card.Suit
}

func (b Bet) Passed(value card.Score) bool {
	return value.Value() >= b.amount*10
}

// Func (b Bet) IsFromTeam(t Team) bool {
//	panic("not implemented")
//	//return b.teamID == t.ID()
// }.

const maxBetAmount = 50

func NewBet(teamID team.ID, amount int, suit card.Suit) Bet {
	if amount < 0 || amount > maxBetAmount {
		panic("invalid bet amount")
	}

	return Bet{teamID, amount, suit}
}
