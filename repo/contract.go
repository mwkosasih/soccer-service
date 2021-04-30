package repo

import (
	"context"

	"github.com/mwkosasih/soccer-service/entity"
)

type TeamRepo interface {
	CreateTeam(ctx context.Context, req entity.Team) error
	GetTeamByID(ctx context.Context, teamID int32) (entity.Team, error)
}

type PlayerRepo interface {
	CreatePlayer(ctx context.Context, req entity.Player) error
	GetPlayerByID(ctx context.Context, ID int32) (entity.Player, error)
	GetPlayerByTeamID(ctx context.Context, teamID int32) ([]entity.Player, error)
}
