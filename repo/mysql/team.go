package mysql

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/mwkosasih/soccer-service/entity"
)

type Team struct {
}

func (l *Team) CreateTeam(ctx context.Context, req entity.Team) error {
	var args []interface{}
	args = append(args, req.Name)
	q := QueryCreateTeam
	_, err := sqlxDB.Exec(sqlxDB.Rebind(q), args...)
	if err != nil {
		log.Errorf("CreateTeam.sqlxDB.Exec query failed: %v", err)
		return err
	}

	return nil
}

func (l *Team) GetTeamByID(ctx context.Context, teamID int32) (entity.Team, error) {
	team := entity.Team{}
	q := QuerySelectTeamByID
	err := sqlxDB.Get(&team, q, teamID)
	if err != nil {
		log.Errorf("GetTeamByID.sqlxDB.Select query failed: %v", err)
		return team, err
	}

	return team, nil
}

func NewTeam() *Team {
	return &Team{}
}
