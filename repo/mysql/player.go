package mysql

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/mwkosasih/soccer-service/entity"
)

type Player struct {
}

func (l *Player) CreatePlayer(ctx context.Context, req entity.Player) error {
	var args []interface{}
	args = append(args, req.TeamID, req.Name, req.Position, req.Height, req.Weight)
	q := QueryCreatePlayer
	_, err := sqlxDB.Exec(sqlxDB.Rebind(q), args...)
	if err != nil {
		log.Errorf("CreatePlayer.sqlxDB.Exec query failed: %v", err)
		return err
	}

	return nil
}

func (l *Player) GetPlayerByID(ctx context.Context, ID int32) (entity.Player, error) {
	player := entity.Player{}
	q := QuerySelectPlayerByID
	err := sqlxDB.Get(&player, q, ID)
	if err != nil {
		log.Errorf("GetPlayerByID.sqlxDB.Select query failed: %v", err)
		return player, err
	}

	return player, nil
}

func (l *Player) GetPlayerByTeamID(ctx context.Context, teamID int32) ([]entity.Player, error) {
	var player []entity.Player
	q := QuerySelectPlayerByTeamID
	err := sqlxDB.Select(&player, q, teamID)
	if err != nil {
		log.Errorf("GetPlayerByTeamID.sqlxDB.Select query failed: %v", err)
		return player, err
	}

	return player, nil
}

func NewPlayer() *Player {
	return &Player{}
}
