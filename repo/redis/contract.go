package redis

import (
	"context"
	"time"
)

var (
	GetTeamKey           = "team:id:%d"
	GetPlayerKey         = "player:id:%d"
	GetPlayerByTeamIDKey = "player:team-id:%d"
	DefaultDuration      = 20 * time.Second
)

type IRedis interface {
	Get(context.Context, string) ([]byte, error)
	Set(context.Context, string, interface{}, time.Duration) error
}
