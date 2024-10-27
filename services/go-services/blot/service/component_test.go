package service

import (
	"blot/internal/blot/ports"
	"blot/internal/common/logging"
	"blot/internal/common/server/grpcserver"
	"blot/internal/common/tests"
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	blotservicepb "blot/internal/common/gen-proto/blotservice/v1beta1"
)

const blotServiceAddr = "localhost:8081"

func TestCreateGameSet(t *testing.T) {
	t.Parallel()

	// ARRANGE
	client := newBlotServiceClient(t)

	// ACT
	resp, err := client.CreateGameSet(context.Background(), &blotservicepb.CreateGameSetRequest{
		Id:         "58F59F84-ADBA-488F-BE2F-4C8F5BC81609",
		PlayerId:   "4FE3928C-8867-4C16-B6DA-8978AE1ABF0E",
		PlayerName: "John",
	})

	// ASERT
	require.NoError(t, err)
	require.NotNil(t, resp)
	re, err := client.GetGameSetForPlayer(context.Background(), &blotservicepb.GetGameSetForPlayerRequest{
		Id:       "58F59F84-ADBA-488F-BE2F-4C8F5BC81609",
		PlayerId: "4FE3928C-8867-4C16-B6DA-8978AE1ABF0E",
	})
	require.NoError(t, err)
	require.NotNil(t, re)
	require.Equal(t, "58f59f84-adba-488f-be2f-4c8f5bc81609", re.GameSet.Id)
	require.Equal(t, "4fe3928c-8867-4c16-b6da-8978ae1abf0e", re.GameSet.Players[0].Id)
	require.Equal(t, "John", re.GameSet.Players[0].Name)
	require.Equal(t, blotservicepb.GameSetStatus_GAME_SET_STATUS_WAITED_FOR_PLAYERS, re.GameSet.Status)
}

func TestJoinGameSet(t *testing.T) {
	t.Parallel()

	// ARRANGE
	client := newBlotServiceClient(t)
	gameSetId := "daeed6ef-e697-4fd1-9748-468e844bd210"
	players := []struct {
		id   string
		name string
	}{
		{"b9ba7e73-d777-4de0-9931-e1ef8a9aa284", "John"},
		{"82607f08-99ac-4b36-b342-fe58cfb0461b", "Jane"},
		{"5f077899-a76c-472f-8357-e892c6da6e54", "Jack"},
		{"97300435-f40f-40ce-9c2e-2866da9b2161", "Jill"},
	}
	_, err := client.CreateGameSet(context.Background(), &blotservicepb.CreateGameSetRequest{
		Id:         gameSetId,
		PlayerId:   players[0].id,
		PlayerName: players[0].name,
	})
	require.NoError(t, err)

	// ACT
	for _, p := range players[1:] {
		_, err = client.JoinGameSet(context.Background(), &blotservicepb.JoinGameSetRequest{
			Id:         gameSetId,
			PlayerId:   p.id,
			PlayerName: p.name,
		})
		require.NoError(t, err)
	}

	// ASERT
	re, err := client.GetGameSetForPlayer(context.Background(), &blotservicepb.GetGameSetForPlayerRequest{
		Id:       gameSetId,
		PlayerId: players[0].id,
	})
	require.NoError(t, err)
	require.NotNil(t, re)
	require.Equal(t, gameSetId, re.GameSet.Id)
	require.Equal(t, blotservicepb.GameSetStatus_GAME_SET_STATUS_READY_TO_START, re.GameSet.Status)
	require.Len(t, re.GameSet.Players, 4)
	for i, p := range players {
		require.Equal(t, p.id, re.GameSet.Players[i].Id)
		require.Equal(t, p.name, re.GameSet.Players[i].Name)
	}
}

func TestStartGameSet(t *testing.T) {
	// ARRANGE
	t.Parallel()
	client := newBlotServiceClient(t)
	gameSetID := "aaeed6ef-e697-4fd1-9748-468e844bd212"
	firstPlayerID := "c9ba7e73-d777-4de0-9931-e1ef8a9aa288"
	prepareGameSetToStart(t, client, gameSetID, firstPlayerID)

	// ACT
	gameId := "b4f8b558-8576-4af2-be01-23f8a8b270b2"
	_, err := client.StartGame(context.Background(), &blotservicepb.StartGameRequest{
		GameSetId: gameSetID,
		GameId:    gameId,
		PlayerId:  firstPlayerID,
	})

	// ASERT
	require.NoError(t, err)
	re, err := client.GetGameSetForPlayer(context.Background(), &blotservicepb.GetGameSetForPlayerRequest{
		Id:       gameSetID,
		PlayerId: firstPlayerID,
	})
	require.NoError(t, err)
	require.NotNil(t, re)
	require.Equal(t, gameSetID, re.GameSet.Id)
	require.Equal(t, blotservicepb.GameSetStatus_GAME_SET_STATUS_PLAYING, re.GameSet.Status)
	require.Equal(t, gameId, re.GameSet.Game.Id)
	require.Equal(t, blotservicepb.GameStatus_GAME_STATUS_BETTING, re.GameSet.Game.Status)
	require.Equal(t, firstPlayerID, re.GameSet.Game.PlayerStates[0].Id)
	require.Len(t, re.GameSet.Game.PlayerStates, 4)
	require.Len(t, re.GameSet.Game.PlayerStates[0].HandCards, 8)
}

func prepareGameSetToStart(t *testing.T, client blotservicepb.BlotServiceClient, setID string, firstPlayerID string) {
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
}

func newBlotServiceClient(t *testing.T) blotservicepb.BlotServiceClient {
	t.Helper()
	conn, err := grpc.NewClient(blotServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return blotservicepb.NewBlotServiceClient(conn)
}

func startService() bool {
	app := NewApplication(context.Background())
	logger := logging.NewLogger(os.Stdout, true, slog.LevelDebug)
	slog.SetDefault(logger)
	go func() {
		err := grpcserver.RunServerOnAddr(
			blotServiceAddr,
			func(server *grpc.Server) {
				svc := ports.NewGrpcServer(app)
				reflection.Register(server)
				blotservicepb.RegisterBlotServiceServer(server, svc)
			})
		if err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
	}()

	ok := tests.WaitForPort(blotServiceAddr)
	if !ok {
		log.Println("Timed out waiting for trainings HTTP to come up")
	}

	return ok
}

func TestMain(m *testing.M) {
	if !startService() {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
