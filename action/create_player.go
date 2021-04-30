package action

import (
	"context"
	"database/sql"

	"github.com/mwkosasih/soccer-service/entity"
	"github.com/mwkosasih/soccer-service/util"
	log "github.com/sirupsen/logrus"

	"github.com/mwkosasih/soccer-service/repo"
)

type CreatePlayer struct {
	playerRepo repo.PlayerRepo
	teamRepo   repo.TeamRepo
}

func (g *CreatePlayer) Handle(ctx context.Context, req entity.Player) error {
	log.Infof("Handle.CreatePlayer request: %v", req)
	// find team by id
	_, err := g.teamRepo.GetTeamByID(ctx, req.TeamID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warnf("CreatePlayer.GetTeamByID with id: %d not found", req.TeamID)
			return util.NotFoundError("team")
		}
		log.Errorf("CreatePlayer.GetTeamByID Failed: %v", err)
		return err
	}

	// create data to `players` table
	err = g.playerRepo.CreatePlayer(ctx, req)
	if err != nil {
		log.Errorf("CreatePlayer to database Failed: %v", err)
		return err
	}

	log.Infof("Handle.CreatePlayer finish")
	return nil
}

func NewCreatePlayer() *CreatePlayer {
	return &CreatePlayer{
		playerRepo: repo.NewPlayer(),
		teamRepo:   repo.NewTeam(),
	}
}
