package action

import (
	"context"
	"log"
	"testing"

	"github.com/mwkosasih/soccer-service/repo/mysql"
	"github.com/mwkosasih/soccer-service/repo/redis"
	"github.com/mwkosasih/soccer-service/util"
	"github.com/stretchr/testify/assert"
)

func actionGetPlayerTest() *GetPlayer {
	// init env
	util.InitEnvTest()

	// connect redis
	redis.Connect()

	// connect mysql
	mysql.Connect()

	return NewGetPlayer()
}

func TestAction_GetPlayer_Success(t *testing.T) {
	app := actionGetPlayerTest()

	playerResp, err := app.Handle(context.TODO(), 1)
	log.Printf("%+v", err)
	log.Printf("%+v", playerResp)
	assert.NoErrorf(t, err, "This should have no error")
}
