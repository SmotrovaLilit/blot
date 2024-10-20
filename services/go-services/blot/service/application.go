package service

import (
	"context"

	"blot/internal/blot/adapters"
	"blot/internal/blot/app"
	"blot/internal/blot/app/command"
	"blot/internal/blot/app/query"
)

func NewApplication(ctx context.Context) app.Application {
	gameSetRepository := adapters.NewGameSetMemoryRepository()

	return app.Application{
		Commands: app.Commands{
			CreateGameSet: command.NewCreateGameSetHandler(gameSetRepository),
			StartNewGame:  command.NewStartNewGameHandler(gameSetRepository),
			PlayCard:      command.NewPlayCardHandler(gameSetRepository),
			JoinGameSet:   command.NewJoinGameSetHandler(gameSetRepository),
			LeaveGameSet:  command.NewLeaveGameSetHandler(gameSetRepository),
		},
		Queries: app.Queries{
			GameSetForPlayer:  query.NewGameSetForPlayerQueryHandler(gameSetRepository),
			GameSetsForPlayer: query.NewGameSetsForPlayerQueryHandler(gameSetRepository),
		},
	}
}
