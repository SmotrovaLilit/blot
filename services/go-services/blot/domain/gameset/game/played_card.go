package game

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/player"
)

type PlayerCard struct {
	playerID player.ID
	card     card.Card
}

func (c PlayerCard) String() string {
	return c.playerID.String() + ":" + c.card.String()
}

func NewPlayerCard(playerID player.ID, card card.Card) PlayerCard {
	if playerID.IsZero() || card.IsZero() {
		panic("invalid arguments, create objects using constructors")
	}
	return PlayerCard{
		playerID: playerID,
		card:     card,
	}
}

func (c PlayerCard) PlayerID() player.ID {
	return c.playerID
}

func (c PlayerCard) Card() card.Card {
	return c.card
}
