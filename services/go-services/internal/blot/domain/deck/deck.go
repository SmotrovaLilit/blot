package deck

import (
	"math/rand"

	"blot/internal/blot/domain/card"
	//"blot/internal/blot/domain/game/player"
)

type Deck struct {
	cards [32]card.Card
}

func NewDeck() Deck {
	deck := Deck{}
	i := 0
	for _, suit := range card.Suits {
		for _, value := range card.Ranks {
			deck.cards[i] = card.NewCard(value, suit)
			i++
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(32, func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

//func (d *Deck) DealCards(players []player.Player) {
//	// TODO может ли он менять чужой агрегат?
//	for i := 0; i < 8; i++ {
//		for j := 0; j < 4; j++ {
//			players[j].AddHandCard(d.cards[i*4+j])
//		}
//	}
//}
