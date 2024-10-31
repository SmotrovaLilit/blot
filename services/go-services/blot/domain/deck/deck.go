package deck

import (
	"math/rand/v2"

	"blot/internal/blot/domain/card"
	// "blot/internal/blot/domain/game/player".
)

type Deck struct {
	cards [32]card.Card
	rand  *rand.Rand
}

func NewDeck(randSource rand.Source) *Deck {
	deck := Deck{
		// nolint
		rand: rand.New(randSource),
	}
	i := 0
	for _, suit := range card.Suits {
		for _, value := range card.Ranks {
			deck.cards[i] = card.NewCard(value, suit)
			i++
		}
	}
	deck.shuffle()
	return &deck
}

func (d *Deck) shuffle() {
	// TODO think about domain purity and side effects
	// https://enterprisecraftsmanship.com/posts/domain-model-purity-completeness/
	// nolint
	d.rand.Shuffle(32, func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) DealCards() [4][]card.Card {
	hands := [4][]card.Card{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			hands[j] = append(hands[j], d.cards[i*4+j])
		}
	}
	return hands
}
