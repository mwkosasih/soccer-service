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

func actionGetTeamTest() *GetTeam {
	// init env
	util.InitEnvTest()

	// connect redis
	redis.Connect()

	// connect mysql
	mysql.Connect()

	return NewGetTeam()
}

func TestAction_GetTeam_Success(t *testing.T) {
	app := actionGetTeamTest()

	teamResp, playerResp, err := app.Handle(context.TODO(), 1)
	log.Printf("%+v", err)
	log.Printf("%+v", teamResp)
	log.Printf("%+v", playerResp)
	assert.NoErrorf(t, err, "This should have no error")
}
