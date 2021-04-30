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

type Player struct {
	player mysql.IPlayer
	redis  redis.IRedis
}

func (l *Player) CreatePlayer(ctx context.Context, req entity.Player) error {
	err := l.player.CreatePlayer(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (l *Player) GetPlayerByID(ctx context.Context, ID int32) (entity.Player, error) {
	// get list from redis
	key := fmt.Sprintf(redis.GetPlayerKey, ID)
	v, err := l.redis.Get(ctx, key)
	if err != nil {
		log.Warnf("Player.RedisGet GetPlayerByID with ID: %d failed", ID)

		// get from database
		resp, err := l.player.GetPlayerByID(ctx, ID)
		if err != nil {
			log.Errorf("Player.GetPlayerByID from DB with ID: %d failed", ID)
			return entity.Player{}, err
		}

		// set redis data
		bytes, _ := json.Marshal(resp)
		key := fmt.Sprintf(redis.GetPlayerKey, ID)
		l.redis.Set(ctx, key, bytes, redis.DefaultDuration)

		return resp, err
	}

	var resp entity.Player
	err = json.Unmarshal(v, &resp)
	if err != nil {
		return entity.Player{}, err
	}

	return resp, nil
}

func (l *Player) GetPlayerByTeamID(ctx context.Context, teamID int32) ([]entity.Player, error) {
	// get list from redis
	key := fmt.Sprintf(redis.GetPlayerByTeamIDKey, teamID)
	v, err := l.redis.Get(ctx, key)
	if err != nil {
		log.Warnf("Player.RedisGet GetPlayerByTeamID with teamID: %d failed", teamID)

		// get from database
		resp, err := l.player.GetPlayerByTeamID(ctx, teamID)
		if err != nil {
			log.Errorf("Player.GetPlayerByTeamID from DB with teamID: %d failed", teamID)
			return []entity.Player{}, err
		}

		// set redis data
		bytes, _ := json.Marshal(resp)
		key := fmt.Sprintf(redis.GetPlayerByTeamIDKey, teamID)
		l.redis.Set(ctx, key, bytes, redis.DefaultDuration)

		return resp, err
	}

	var resp []entity.Player
	err = json.Unmarshal(v, &resp)
	if err != nil {
		return []entity.Player{}, err
	}

	return resp, nil
}

func NewPlayer() *Player {
	return &Player{
		player: mysql.NewPlayer(),
		redis:  redis.NewRedis(),
	}
}
