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

func (s Suit) String() string {
	return s.value
}

func NewSuit(suitString string) Suit {
	for _, suit := range Suits {
		if suit.value == suitString {
			return suit
		}
	}
	panic(fmt.Sprintf("Invalid suit: %s", suitString))
}
