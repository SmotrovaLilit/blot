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

func (c Card) IsStronger(card Card, trump Suit) bool {
	panic("not implemented")
}

func NewCard(rank Rank, suit Suit) Card {
	return Card{rank, suit}
}
