package action

import (
	"context"
	"github.com/mwkosasih/soccer-service/entity"
	"log"
	"testing"

	"github.com/mwkosasih/soccer-service/repo/mysql"
	"github.com/mwkosasih/soccer-service/repo/redis"
	"github.com/mwkosasih/soccer-service/util"
	"github.com/stretchr/testify/assert"
)

func actionCreatePlayerTest() *CreatePlayer {
	// init env
	util.InitEnvTest()

	// connect redis
	redis.Connect()

	// connect mysql
	mysql.Connect()

	return NewCreatePlayer()
}

func TestAction_CreatePlayer_Success(t *testing.T) {
	app := actionCreatePlayerTest()

	req := entity.Player{
		TeamID:   1,
		Name:     "Wildan",
		Position: "Kiper",
		Height:   190,
		Weight:   30,
	}

	err := app.Handle(context.TODO(), req)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}
