package game

import (
	"errors"

	"blot/internal/blot/domain/card"
)

var ErrTableFull = errors.New("table is full")

type TableCards struct {
	cards []PlayerCard
}

const maxCardsInRound = 4

func NewTableCards(cards []PlayerCard) TableCards {
	if len(cards) > maxCardsInRound {
		panic("too many cards")
	}
	return TableCards{
		cards: cards,
	}
}

func (t TableCards) String() string {
	str := ""
	for _, c := range t.cards {
		str += c.String() + " "
	}
	return str
}

func (t TableCards) Len() int {
	return len(t.cards)
}

func (t TableCards) IsFull() bool {
	return t.Len() == maxCardsInRound
}

func (t TableCards) Add(c PlayerCard) (TableCards, error) {
	if t.IsFull() {
		return t, ErrTableFull
	}
	newTable := t.Clone()
	newTable.cards = append(newTable.cards, c)
	return newTable, nil
}

func (t TableCards) Clone() TableCards {
	cards := make([]PlayerCard, len(t.cards))
	copy(cards, t.cards)
	return TableCards{
		cards: cards,
	}
}

func (t TableCards) CalculateWinner(trump card.Suit) PlayerCard {
	if !t.IsFull() {
		return PlayerCard{}
	}
	winner := t.cards[0]
	for _, c := range t.cards[1:] {
		if c.card.Beats(winner.card, trump) {
			winner = c
		}
	}
	return winner
}

func (t TableCards) CalculateScore(trump card.Suit) card.Score {
	score := card.NewScore(0)
	for _, c := range t.cards {
		score = score.Add(c.card.GetScore(trump))
	}
	return score
}

func (t TableCards) HasCard(argsCard card.Card) bool {
	for _, c := range t.cards {
		if c.card.Equal(argsCard) {
			return true
		}
	}
	return false
}

func (t TableCards) Cards() []PlayerCard {
	cards := make([]PlayerCard, len(t.cards))
	copy(cards, t.cards)
	return cards
}
