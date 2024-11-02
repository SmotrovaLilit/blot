package game

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/player"
)

type ErrCardNotFound struct {
	PlayerID string
	Card     string
}

func (e ErrCardNotFound) Error() string {
	return "card not found in player's hand: " + e.Card + " for player: " + e.PlayerID
}

type PlayerState struct {
	id        player.ID
	handCards []card.Card
}

func NewPlayerState(id player.ID, handCards []card.Card) *PlayerState {
	return &PlayerState{id, handCards}
}

func (s PlayerState) ID() player.ID {
	return s.id
}

func (s PlayerState) HandCards() []card.Card {
	return s.handCards
}

func (s *PlayerState) Clone() PlayerState {
	cards := make([]card.Card, len(s.handCards))
	copy(cards, s.handCards)
	return PlayerState{s.id, cards}
}

func (s *PlayerState) RemoveCard(c card.Card) error {
	for i, ca := range s.handCards {
		if ca.Equal(c) {
			s.handCards = append(s.handCards[:i], s.handCards[i+1:]...)
			return nil
		}
	}
	return ErrCardNotFound{Card: c.String(), PlayerID: s.id.String()}
}

func (s *PlayerState) CanRemoveCard(c card.Card) error {
	for _, ca := range s.handCards {
		if ca.Equal(c) {
			return nil
		}
	}
	return ErrCardNotFound{Card: c.String(), PlayerID: s.id.String()}
}

func UnmarshalFromDatabasePlayerState(id player.ID, cards []card.Card) PlayerState {
	return PlayerState{id, cards}
}
