package tests

import (
	"math/rand/v2"
	"testing"

	"blot/internal/blot/domain/gameset/game/bet"

	"github.com/stretchr/testify/require"

	"blot/internal/blot/domain/card"
	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/gameset/game"
	"blot/internal/blot/domain/gameset/player"
)

func PrepareGameSetToStartGame(t *testing.T) *gameset.GameSet {
	t.Helper()
	firstPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a04")
	secondPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a05")
	thirdPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a06")
	fourthPlayerID := player.MustNewID("4eb00c05-7f64-47b0-81bf-d0977bff0a07")
	set := gameset.NewGameSet(
		gameset.MustNewID("317c8f91-14ef-4582-aaa0-636b5d2ca0c2"),
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
	require.Equal(t, gameset.StatusReadyToStart, set.Status())
	return set
}

func PrepareGameSetToPlayCard(t *testing.T) *gameset.GameSet {
	t.Helper()
	set := PrepareGameSetToSetBet(t)
	set.MustSetBet(set.OwnerID(), card.SuitSpades, bet.MustNewAmount(8))
	require.Equal(t, gameset.StatusPlaying, set.Status())
	newGame := set.LastGame()
	lastRound, err := newGame.LastRound()
	require.NoError(t, err)
	require.Equal(t, game.RoundNumber1, lastRound.Number())
	return set
}

func PrepareGameSetToSetBet(t *testing.T) *gameset.GameSet {
	t.Helper()
	set := PrepareGameSetToStartGame(t)
	set.MustStartGame(
		game.MustNewID("937cc314-7cf3-4918-8c16-f1699eee89d9"),
		set.OwnerID(),
		// nolint
		rand.New(rand.NewPCG(0, 0)),
	)
	g := set.LastGame()
	require.Equal(t, gameset.StatusPlaying, set.Status())
	require.Equal(t, game.StatusBetting, g.Status())
	return set
}
