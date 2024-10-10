package app

import "blot/internal/blot/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateGameSet command.CreateGameSetHandler
	StartNewGame  command.StartNewGameHandler
	PlayCard      command.PlayCardHandler
}
