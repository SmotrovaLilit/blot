package service

import (
	"context"

	"blot/internal/blot/adapters"
	"blot/internal/blot/app"
	"blot/internal/blot/app/command"
)

func NewApplication(ctx context.Context) app.Application {
	gameSetRepository := adapters.NewGameSetMemoryRepository()
	return app.Application{
		Commands: app.Commands{
			CreateGameSet: command.NewCreateGameSetHandler(gameSetRepository),
			StartNewGame:  command.NewStartNewGameHandler(gameSetRepository),
			PlayCard:      command.NewPlayCardHandler(gameSetRepository),
		},
	}
}
