package gameset

import "blot/internal/blot/domain/card"

type Turn struct {
	roundNumber RoundNumber
	//player      *Player
	card card.Card
}
