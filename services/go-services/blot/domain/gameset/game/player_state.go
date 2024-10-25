package game

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/player"
)

type PlayerState struct {
	id        player.ID
	handCards []card.Card
}

func NewPlayerState(id player.ID, handCards []card.Card) PlayerState {
	return PlayerState{id, handCards}
}

func (s PlayerState) ID() player.ID {
	return s.id
}

func (s PlayerState) HandCards() []card.Card {
	return s.handCards
}

func UnmarshalFromDatabasePlayerState(id player.ID, cards []card.Card) PlayerState {
	return PlayerState{id, cards}
}
