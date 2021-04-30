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

func actionCreateTeamTest() *CreateTeam {
	// init env
	util.InitEnvTest()

	// connect redis
	redis.Connect()

	// connect mysql
	mysql.Connect()

	return NewCreateTeam()
}

func TestAction_CreateTeam_Success(t *testing.T) {
	app := actionCreateTeamTest()

	err := app.Handle(context.TODO(), "Manchester United")
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}
