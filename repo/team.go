package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mwkosasih/soccer-service/repo/redis"
	log "github.com/sirupsen/logrus"

	"github.com/mwkosasih/soccer-service/entity"
	"github.com/mwkosasih/soccer-service/repo/mysql"
)

type Team struct {
	team  mysql.ITeam
	redis redis.IRedis
}

func (l *Team) CreateTeam(ctx context.Context, req entity.Team) error {
	err := l.team.CreateTeam(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (l *Team) GetTeamByID(ctx context.Context, teamID int32) (entity.Team, error) {
	// get list from redis
	key := fmt.Sprintf(redis.GetTeamKey, teamID)
	v, err := l.redis.Get(ctx, key)
	if err != nil {
		log.Warnf("Team.RedisGet GetTeamByID with ID: %d failed", teamID)

		// get from database
		resp, err := l.team.GetTeamByID(ctx, teamID)
		if err != nil {
			log.Errorf("Team.GetTeamByID from DB with ID: %d failed", teamID)
			return entity.Team{}, err
		}

		// set redis data
		bytes, _ := json.Marshal(resp)
		key := fmt.Sprintf(redis.GetTeamKey, teamID)
		l.redis.Set(ctx, key, bytes, redis.DefaultDuration)

		return resp, err
	}

	var resp entity.Team
	err = json.Unmarshal(v, &resp)
	if err != nil {
		return entity.Team{}, err
	}

	return resp, nil
}

func NewTeam() *Team {
	return &Team{
		team:  mysql.NewTeam(),
		redis: redis.NewRedis(),
	}
}
