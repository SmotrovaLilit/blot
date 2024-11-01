package gameset

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/bet"
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/player"
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"testing"
)

func TestPlayingGame(t *testing.T) {
	t.Parallel()
	firstPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04")
	secondPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a05")
	thirdPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a06")
	fourthPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a07")
	set := NewGameSet(
		MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2"),
		player.New(
			firstPlayerID,
			player.MustNewName("John"),
		),
	)
	set.MustJoin(player.New(
		secondPlayerID,
		player.MustNewName("Jane"),
	))
	set.MustJoin(player.New(
		thirdPlayerID,
		player.MustNewName("Jack"),
	))
	set.MustJoin(player.New(
		fourthPlayerID,
		player.MustNewName("Jill"),
	))
	set.MustStartGame(
		game.MustNewID("937cc314-7cf3-4918-8c16-f1699eee89d9"),
		firstPlayerID,
		rand.NewPCG(0, 0),
	)

	g := set.LastGame()
	require.Equal(t, game.StatusBetting, g.Status())
	err := set.SetBet(firstPlayerID, card.SuitSpades, bet.MustNewAmount(8))
	require.NoError(t, err)

	g = set.LastGame()
	require.Equal(t, game.StatusPlaying, g.Status())
	pCard := g.MustPlayerState(firstPlayerID).HandCards()[0]
	err = set.PlayCard(firstPlayerID, pCard)
	require.NoError(t, err)
	g = set.LastGame()
	state := g.FirstPlayerState()
	require.Len(t, state.HandCards(), 7)
	requreNotContainsCard(t, g.FirstPlayerState().HandCards(), pCard)
	// TODO check table cards contain the played card
	// TODO check next player is the next player
	// TODO может аграгат должен быть как фасад, мы не должны обращаться к внутренним сущностям напрямую? или это нормально?
}

func requreNotContainsCard(t *testing.T, cards []card.Card, pCard card.Card) {
	t.Helper()
	for _, c := range cards {
		if c.Equal(pCard) {
			require.Fail(t, "card found in player's hand")
		}
	}
}
