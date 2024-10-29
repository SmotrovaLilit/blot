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

func NewCard(rank Rank, suit Suit) Card {
	return Card{rank, suit}
}

func UnmarshalFromDatabase(rank Rank, suit Suit) Card {
	return Card{rank, suit}
}
