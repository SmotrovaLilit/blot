package card

type Card struct {
	rank Rank
	suit Suit
}

func (c Card) GetScore(trump Suit) Score {
	if c.suit == trump {
		return c.rank.GetTrumpScore()
	}
	return c.rank.GetScore()
}

func (c Card) IsStronger(_ Card, _ Suit) bool {
	panic("not implemented")
}

func (c Card) Suit() Suit {
	return c.suit
}

func (c Card) Rank() Rank {
	return c.rank
}

func (c Card) String() string {
	return c.rank.String() + " of " + c.suit.String()
}

func (c Card) Equal(c2 Card) bool {
	return c.rank == c2.rank && c.suit == c2.suit
}

func (c Card) IsZero() bool {
	return c.rank.IsZero() && c.suit.IsZero()
}

func (c Card) Beats(winner Card, trump Suit) bool {
	if c.suit.Equal(winner.suit) {
		return c.rank.Beats(winner.rank, trump.Equal(c.suit))
	}
	if c.suit.Equal(trump) {
		return true
	}
	return false
}

func NewCard(rank Rank, suit Suit) Card {
	return Card{rank, suit}
}

func UnmarshalFromDatabase(rank Rank, suit Suit) Card {
	return Card{rank, suit}
}
