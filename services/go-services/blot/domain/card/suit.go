package card

import (
	"fmt"
)

var (
	SuitClubs    = Suit{"Clubs"}
	SuitDiamonds = Suit{"Diamonds"}
	SuitHearts   = Suit{"Hearts"}
	SuitSpades   = Suit{"Spades"}

	Suits = []Suit{SuitClubs, SuitDiamonds, SuitHearts, SuitSpades}
)

type Suit struct {
	value string
}

func NewSuit(suitString string) Suit {
	for _, suit := range Suits {
		if suit.value == suitString {
			return suit
		}
	}
	panic(fmt.Sprintf("Invalid suit: %s", suitString))
}

func (s Suit) String() string {
	return s.value
}

func (s Suit) IsZero() bool {
	return Suit{} == s
}

func (s Suit) Equal(s2 Suit) bool {
	return s.value == s2.value
}
