package old

import "blot/internal/blot/domain/card"

type PlayedCard struct {
	card   card.Card
	Player Player
}

func NewPlayedCard(c card.Card, p Player) PlayedCard {
	return PlayedCard{c, p}
}
