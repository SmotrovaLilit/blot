package game

import "blot/internal/blot/domain/card"

type Round struct {
	number RoundNumber
	table  TableCards // TODO maybe use one value object for both table and rounds
}

type ErrCannotPlayCard struct {
	Card  string
	Table string
}

func (e ErrCannotPlayCard) Error() string {
	return "cannot play card " + e.Card + " on table " + e.Table
}

func NewRound(number RoundNumber, table TableCards) Round {
	return Round{
		number: number,
		table:  table,
	}
}

func (r Round) Number() RoundNumber {
	return r.number
}

func (r Round) Table() TableCards {
	return r.table
}

func (r Round) PlayCard(c PlayerCard) (Round, error) {
	if err := canPlayCard(r.table, c); err != nil {
		return Round{}, err
	}
	newTable, err := r.table.Add(c)
	if err != nil {
		return Round{}, err
	}
	return NewRound(r.number, newTable), nil
}

func (r Round) String() string {
	return r.Number().String() + ":" + r.table.String()
}

func (r Round) Finished() bool {
	return r.table.IsFull()
}

func canPlayCard(table TableCards, playedCard PlayerCard) error {
	if !table.IsFull() {
		return nil
	}
	// TODO: implement this
	return ErrCannotPlayCard{Card: playedCard.String(), Table: table.String()} // TODO
}

func (r *Round) Clone() Round {
	return NewRound(r.number, r.table.Clone())
}

func (r *Round) CalculateWinner(trump card.Suit) PlayerCard {
	return r.table.CalculateWinner(trump)
}

func (r Round) CalculateScore(trump card.Suit) card.Score {
	return r.table.CalculateScore(trump)
}

func (r Round) HasCard(argsCard card.Card) bool {
	return r.table.HasCard(argsCard)
}

func (r Round) TableCards() []PlayerCard {
	return r.table.Cards()
}
