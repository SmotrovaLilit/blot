package service

import (
	"context"
	"os"
	"runtime/debug"

	"blot/internal/blot/app/command/creategameset"

	"blot/internal/blot/adapters"
	"blot/internal/blot/app"
	"blot/internal/blot/app/command"
	"blot/internal/blot/app/query"
)

func NewApplication(ctx context.Context) app.Application {
	gameSetRepository := adapters.NewGameSetMemoryRepository()

	// We can move it to common adapters, to reuse logic for all services
	buildInfo, _ := debug.ReadBuildInfo()
	pid := os.Getpid()

	return app.Application{
		Commands: app.Commands{
			CreateGameSet: creategameset.NewHandler(gameSetRepository),
			StartGame:     command.NewStartGameHandler(gameSetRepository),
			PlayCard:      command.NewPlayCardHandler(gameSetRepository),
			JoinGameSet:   command.NewJoinGameSetHandler(gameSetRepository),
			LeaveGameSet:  command.NewLeaveGameSetHandler(gameSetRepository),
		},
		Queries: app.Queries{
			GameSetForPlayer:  query.NewGameSetForPlayerQueryHandler(gameSetRepository),
			GameSetsForPlayer: query.NewGameSetsForPlayerQueryHandler(gameSetRepository),
		},
		Info: app.Info{
			Name:      "blot",    // TODO got it from config
			Version:   "v1beta1", // TODO got it from config
			BuildPath: buildInfo.Main.Path,
			GoVersion: buildInfo.GoVersion,
			PID:       pid,
		},
	}
}
