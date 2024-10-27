package app

import (
	"context"
	"log/slog"

	"blot/internal/blot/app/command"
	"blot/internal/blot/app/command/creategameset"
	"blot/internal/blot/app/query"
	"blot/internal/common/logging"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Info     Info
}

type Info struct {
	Name      string
	Version   string
	BuildPath string
	GoVersion string
	PID       int
}

type Commands struct {
	CreateGameSet creategameset.Handler
	PlayCard      command.PlayCardHandler
	JoinGameSet   command.JoinGameSetHandler
	LeaveGameSet  command.LeaveGameSetHandler
	StartGame     command.StartGameHandler
}

type Queries struct {
	GameSetForPlayer  query.GameSetForPlayerQueryHandler
	GameSetsForPlayer query.GameSetsForPlayerQueryHandler
}

func (ap Application) AppendCtxWithApplicationLoggingFields(ctx context.Context) context.Context {
	// TODO move this logic in common, to avoid repeating kye names in all services
	return logging.AppendCtx(
		ctx,
		slog.Any("app", ap.Info.Name),
		slog.Any("app_version", ap.Info.Version),
		slog.Any("build_path", ap.Info.BuildPath),
		slog.Any("go_version", ap.Info.GoVersion),
	)
}
