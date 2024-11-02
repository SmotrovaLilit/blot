package bet

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/team"
)

type Bet struct {
	teamID team.ID
	amount Amount
	suit   card.Suit
}

// Func (b Bet) Passed(value card.Score) bool {
//	return value.Value() >= b.amount*10
// }.

// Func (b Bet) IsFromTeam(t Team) bool {
//	panic("not implemented")
//	//return b.teamID == t.ID()
// }.

func (b Bet) Trump() card.Suit {
	return b.suit
}

func (b Bet) Amount() Amount {
	return b.amount
}

func NewBet(teamID team.ID, amount Amount, trump card.Suit) Bet {
	if amount.IsZero() || trump.IsZero() || teamID.IsZero() {
		panic("invalid arguments, create objects using constructors")
	}
	return Bet{teamID, amount, trump}
}

func (b Bet) IsZero() bool {
	return b == Bet{}
}
