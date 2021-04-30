package action

import (
	"context"
	"github.com/mwkosasih/soccer-service/entity"
	log "github.com/sirupsen/logrus"

	"github.com/mwkosasih/soccer-service/repo"
)

type CreateTeam struct {
	teamRepo   repo.TeamRepo
}

func (g *CreateTeam) Handle(ctx context.Context, teamName string) error {
	log.Infof("Handle.CreateTeam request teamName: %s", teamName)
	// create data to `teams` table
	createTeamReq := entity.Team{
		Name: teamName,
	}
	err := g.teamRepo.CreateTeam(ctx, createTeamReq)
	if err != nil {
		log.Errorf("CreateTeam to database Failed: %v", err)
		return err
	}

	log.Infof("Handle.CreateTeam finish")
	return nil
}

func NewCreateTeam() *CreateTeam {
	return &CreateTeam{
		teamRepo: repo.NewTeam(),
	}
}
