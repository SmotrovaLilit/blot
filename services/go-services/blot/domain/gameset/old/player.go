package old

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/team"
	"blot/internal/blot/domain/user"
)

type Player struct {
	userID       user.ID
	teamID       team.ID
	handCards    []card.Card
	discardCards []card.Card
}

func NewPlayer(user user.ID, teamID team.ID) Player {
	return Player{
		userID: user, handCards: []card.Card{}, discardCards: []card.Card{}, teamID: teamID,
	}
}

func (p Player) AddHandCard(c card.Card) {
	p.handCards = append(p.handCards, c)
}

func (p Player) AddDiscardCards(trick [4]card.Card) {
	p.discardCards = append(p.discardCards, trick[:]...)
}

func (p Player) PlayCard(card card.Card) {
	for i, c := range p.handCards {
		if c == card {
			// remove card from hand
			p.handCards = append(p.handCards[:i], p.handCards[i+1:]...)
			return
		}
	}
	panic("card not found")
}

func (p Player) GetHandCards() []card.Card {
	return p.handCards
}

func (p Player) CalculateScore(trump card.Suit) card.Score {
	score := card.NewScore(0)
	for _, c := range p.discardCards {
		score = score.Add(c.GetScore(trump))
	}
	return score
}

func (p Player) Equal(pl Player) bool {
	return p.userID == pl.userID
}

func (p Player) ID() user.ID {
	return p.userID
}
