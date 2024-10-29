package service

import (
	"testing"

	blotservicepb "blot/internal/common/gen-proto/blotservice/v1beta1"
)

func requireLastPlayedCardIs(t *testing.T, _ *blotservicepb.GameSet, _ *blotservicepb.Card) {
	t.Helper()
	panic("not implemented")
	// TODO get last played card from game set
	// lastCard := set.Game
	// require.NotNil(t, lastCard)
	// require.Equal(t, b.Suit, lastCard.Suit)
	// require.Equal(t, b.Rank, lastCard.Rank)
}

func requireNextTurnPlayerIs(t *testing.T, _ *blotservicepb.GameSet, _ string) {
	t.Helper()
	panic("not implemented")
}

func requirePlayerNotContainsCard(
	t *testing.T,
	set *blotservicepb.GameSet,
	playerID string,
	card *blotservicepb.Card,
) {
	t.Helper()
	handCards := make([]*blotservicepb.Card, 0)
	find := false
	for _, p := range set.Game.PlayerStates {
		if p.Id == playerID {
			handCards = p.HandCards
			find = true
			break
		}
	}
	if !find {
		t.Fatalf("player %s not found in gameset", playerID)
	}
	for _, c := range handCards {
		if c.Suit == card.Suit && c.Rank == card.Rank {
			t.Fatalf("card %v is in the list", card)
		}
	}
}

func firstPlayerCard(t *testing.T, set *blotservicepb.GameSet) *blotservicepb.Card {
	t.Helper()
	return set.Game.PlayerStates[0].HandCards[0]
}

func firstPlayerID(t *testing.T, set *blotservicepb.GameSet) string {
	t.Helper()
	return set.Players[0].Id
}
