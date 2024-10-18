package app

import (
	"blot/internal/blot/app/command"
	"blot/internal/blot/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateGameSet command.CreateGameSetHandler
	StartNewGame  command.StartNewGameHandler
	PlayCard      command.PlayCardHandler
}

type Queries struct {
	GameSetForPlayer query.GameSetForPlayerQueryHandler
}
