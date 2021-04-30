package action

import (
	"context"
	"database/sql"

	"github.com/mwkosasih/soccer-service/entity"
	"github.com/mwkosasih/soccer-service/util"
	log "github.com/sirupsen/logrus"

	"github.com/mwkosasih/soccer-service/repo"
)

type GetTeam struct {
	playerRepo repo.PlayerRepo
	teamRepo   repo.TeamRepo
}

func (g *GetTeam) Handle(ctx context.Context, teamID int32) (teamResp entity.Team, playerResp []entity.Player, err error) {
	log.Infof("Handle.GetTeam request id: %d", teamID)

	teamResp, err = g.teamRepo.GetTeamByID(ctx, teamID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warnf("GetTeam.GetTeamByID with id: %d not found", teamID)
			return teamResp, playerResp, util.NotFoundError("team")
		}
		log.Errorf("GetTeam.GetTeamByID Failed: %v", err)
		return teamResp, playerResp, err
	}

	playerResp, err = g.playerRepo.GetPlayerByTeamID(ctx, teamID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warnf("GetTeam.GetPlayerByTeamID with id: %d not found", teamID)
			return teamResp, playerResp, nil
		}
		log.Errorf("GetTeam.GetPlayerByTeamID Failed: %v", err)
		return teamResp, playerResp, err
	}

	log.Infof("Handle.GetTeam finish")
	return teamResp, playerResp, nil
}

func NewGetTeam() *GetTeam {
	return &GetTeam{
		playerRepo: repo.NewPlayer(),
		teamRepo:   repo.NewTeam(),
	}
}
