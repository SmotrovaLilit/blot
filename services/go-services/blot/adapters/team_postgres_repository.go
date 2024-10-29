package adapters

import (
	"context"

	"blot/internal/blot/domain/team"
)

type TeamPostgresRepository struct {
}

func NewTeamPostgresRepository() *TeamPostgresRepository {
	return &TeamPostgresRepository{}
}

func (t TeamPostgresRepository) Get(ctx context.Context, id team.ID) (*team.Team, error) {
	// TODO implement me
	panic("implement me")
}
