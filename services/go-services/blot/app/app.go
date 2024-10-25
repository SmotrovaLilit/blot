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
	PlayCard      command.PlayCardHandler
	JoinGameSet   command.JoinGameSetHandler
	LeaveGameSet  command.LeaveGameSetHandler
	StartGame     command.StartGameHandler
}

type Queries struct {
	GameSetForPlayer  query.GameSetForPlayerQueryHandler
	GameSetsForPlayer query.GameSetsForPlayerQueryHandler
}
