package gameset

import (
	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/game/bet"
	"blot/internal/blot/domain/gameset/player"
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"strconv"
	"testing"
)

func TestPlayingGame(t *testing.T) {
	t.Parallel()
	firstPlayerID := player.MustNewID("0eb00c05-7f64-47b0-81bf-d0977bff0a04")
	secondPlayerID := player.MustNewID("1eb00c05-7f64-47b0-81bf-d0977bff0a05")
	thirdPlayerID := player.MustNewID("2eb00c05-7f64-47b0-81bf-d0977bff0a06")
	fourthPlayerID := player.MustNewID("3eb00c05-7f64-47b0-81bf-d0977bff0a07")
	players := []player.ID{
		firstPlayerID,
		secondPlayerID,
		thirdPlayerID,
		fourthPlayerID,
	}
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

	// Set bet
	g := set.LastGame()
	require.Equal(t, game.StatusBetting, g.Status())
	err := set.SetBet(firstPlayerID, card.SuitSpades, bet.MustNewAmount(8))
	require.NoError(t, err)
	g = set.LastGame()
	require.Equal(t, game.StatusPlaying, g.Status())

	// Play card
	for roundNumber := 1; roundNumber <= 8; roundNumber++ {
		for i := 0; i < 4; i++ {
			t.Run("round "+strconv.Itoa(roundNumber)+" turn "+strconv.Itoa(i+1), func(t *testing.T) {
				playerIndex := (i + roundNumber - 1) % 4
				pCard := g.MustPlayerState(players[playerIndex]).HandCards()[0]
				err = set.PlayCard(players[playerIndex], pCard)
				require.NoError(t, err)
				g = set.LastGame()
				state := g.MustPlayerState(players[playerIndex])
				requreNotContainsCard(t, state.HandCards(), pCard)
				r, err := g.Round(roundNumber)
				require.NoError(t, err)
				require.Equal(t, roundNumber, r.Number().Value())
				require.True(t, r.HasCard(pCard))
			})
		}
	}
	require.Equal(t, game.StatusFinished, g.Status())
	// TODO: check bellow depends on the seed of the random number generator
	g = set.LastGame()
	r, err := g.Round(1)
	require.NoError(t, err)
	require.Equal(t, card.NewCard(card.RankEight, card.SuitSpades), r.CalculateWinner(g.Bet().Trump()).Card(), "trump:"+g.Bet().Trump().String()+", cards: "+r.Table().String())

	r, err = g.Round(2)
	require.NoError(t, err)
	require.Equal(t, card.NewCard(card.RankQueen, card.SuitSpades), r.CalculateWinner(g.Bet().Trump()).Card(), "trump:"+g.Bet().Trump().String()+", cards: "+r.Table().String())

	r, err = g.Round(5)
	require.NoError(t, err)
	require.Equal(t, card.NewCard(card.RankKing, card.SuitHearts), r.CalculateWinner(g.Bet().Trump()).Card(), "trump:"+g.Bet().Trump().String()+", cards: "+r.Table().String())

	r, err = g.Round(8)
	require.NoError(t, err)
	require.Equal(t, card.NewCard(card.RankAce, card.SuitSpades), r.CalculateWinner(g.Bet().Trump()).Card(), "trump:"+g.Bet().Trump().String()+", cards: "+r.Table().String())
}

func requreNotContainsCard(t *testing.T, cards []card.Card, pCard card.Card) {
	t.Helper()
	for _, c := range cards {
		if c.Equal(pCard) {
			require.Fail(t, "card found in player's hand")
		}
	}
}
