package action

import (
	"context"
	"database/sql"

	"github.com/mwkosasih/soccer-service/entity"
	"github.com/mwkosasih/soccer-service/util"
	log "github.com/sirupsen/logrus"

	"github.com/mwkosasih/soccer-service/repo"
)

type GetPlayer struct {
	playerRepo repo.PlayerRepo
}

func (g *GetPlayer) Handle(ctx context.Context, id int32) (resp entity.Player, err error) {
	log.Infof("Handle.GetPlayer request id: %d", id)
	resp, err = g.playerRepo.GetPlayerByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warnf("GetPlayer.GetPlayerByID with id: %d not found", id)
			return resp, util.NotFoundError("player")
		}
		log.Errorf("GetPlayer.GetPlayerByID Failed: %v", err)
		return resp, err
	}

	log.Infof("Handle.GetPlayer finish")
	return resp, nil
}

func NewGetPlayer() *GetPlayer {
	return &GetPlayer{
		playerRepo: repo.NewPlayer(),
	}
}
