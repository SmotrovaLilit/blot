package command

import (
	"blot/internal/common/decorator"
	"context"

	"blot/internal/blot/domain/gameset"
	"blot/internal/blot/domain/team"
)

type CreateGameSet struct {
	TeamID1, TeamID2 team.ID
	ID               gameset.ID
	FirstGameID      gameset.GameID
}

type createGameSetHandler struct {
	gameSetRepository gameset.Repository
	teamRepository    team.Repository
}

type CreateGameSetHandler decorator.CommandHandler[CreateGameSet]

func NewCreateGameSetHandler(gameSetRepository gameset.Repository, teamRepository team.Repository) CreateGameSetHandler {
	if gameSetRepository == nil || teamRepository == nil {
		panic("gameRepository, gameSetRepository, teamRepository cannot be nil")
	}
	return createGameSetHandler{
		gameSetRepository: gameSetRepository,
		teamRepository:    teamRepository,
	}
}

func (h createGameSetHandler) Handle(ctx context.Context, cmd CreateGameSet) error {
	t1, err := h.teamRepository.Get(ctx, cmd.TeamID1)
	if err != nil {
		return err
	}
	t2, err := h.teamRepository.Get(ctx, cmd.TeamID2)
	if err != nil {
		return err
	}
	set, err := gameset.NewGameSet(
		cmd.ID,
		cmd.FirstGameID,
		gameset.NewTeam(cmd.TeamID1, t1.FirstPlayerID(), t1.SecondPlayerID()),
		gameset.NewTeam(cmd.TeamID2, t2.FirstPlayerID(), t2.SecondPlayerID()),
	)
	err = h.gameSetRepository.Create(ctx, set)
	if err != nil {
		return err
	}
	return nil
}
