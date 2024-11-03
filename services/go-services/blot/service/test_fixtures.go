package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	blotservicepb "blot/internal/common/gen-proto/blotservice/v1beta1"
)

func prepareGameSetToStart(t *testing.T, client blotservicepb.BlotServiceClient, setID string, firstPlayerID string) *blotservicepb.GameSet {
	t.Helper()
	_, err := client.CreateGameSet(context.Background(), &blotservicepb.CreateGameSetRequest{
		Id:         setID,
		PlayerId:   firstPlayerID,
		PlayerName: "John",
	})
	require.NoError(t, err)
	players := []struct {
		id   string
		name string
	}{
		{"b9ba7e73-d777-4de0-9931-e1ef8a9aa284", "Jane"},
		{"82607f08-99ac-4b36-b342-fe58cfb0461b", "Jack"},
		{"5f077899-a76c-472f-8357-e892c6da6e54", "Jill"},
	}
	for _, p := range players {
		_, err = client.JoinGameSet(context.Background(), &blotservicepb.JoinGameSetRequest{
			Id:         setID,
			PlayerId:   p.id,
			PlayerName: p.name,
		})
		require.NoError(t, err)
	}
	resp, err := client.GetGameSetForPlayer(context.Background(), &blotservicepb.GetGameSetForPlayerRequest{
		Id:       setID,
		PlayerId: firstPlayerID,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	return resp.GameSet
}

func prepareGameSetToPlayCard(t *testing.T, setID string, client blotservicepb.BlotServiceClient) *blotservicepb.GameSet {
	t.Helper()
	set := prepareGameSetToSetBet(t, setID, client)
	_, err := client.SetBet(context.Background(), &blotservicepb.SetBetRequest{
		GameSetId: set.Id,
		PlayerId:  set.Players[0].Id,
		Trump:     blotservicepb.Suit_SUIT_DIAMONDS,
		Amount:    8,
	})
	require.NoError(t, err)
	resp, err := client.GetGameSetForPlayer(context.Background(), &blotservicepb.GetGameSetForPlayerRequest{
		Id:       set.Id,
		PlayerId: set.Players[0].Id,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	return resp.GameSet
}

func prepareGameSetToSetBet(t *testing.T, setID string, client blotservicepb.BlotServiceClient) *blotservicepb.GameSet {
	t.Helper()
	set := prepareGameSetToStart(t, client, setID, "a9ba7e73-d777-4de0-9931-e1ef8a9aa355")
	_, err := client.StartGame(context.Background(), &blotservicepb.StartGameRequest{
		GameSetId: set.Id,
		GameId:    "b4f8b558-8576-4af2-be01-23f8a8b270b8",
		PlayerId:  set.Players[0].Id,
	})
	require.NoError(t, err)

	resp, err := client.GetGameSetForPlayer(context.Background(), &blotservicepb.GetGameSetForPlayerRequest{
		Id:       set.Id,
		PlayerId: set.Players[0].Id,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	return resp.GameSet
}
