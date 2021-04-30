package repo

import (
	"context"
	"github.com/mwkosasih/soccer-service/entity"
	"github.com/mwkosasih/soccer-service/repo/mysql"
	"github.com/mwkosasih/soccer-service/repo/redis"
	"github.com/mwkosasih/soccer-service/util"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func teamRepoMysqlTest() TeamRepo {
	// init env
	util.InitEnvTest()

	// connect redis
	redis.Connect()

	// connect mysql
	mysql.Connect()

	return NewTeam()
}

func TestMysql_CreateTeam_Success(t *testing.T) {
	app := teamRepoMysqlTest()

	req := entity.Team{
		Name: "Chelsea",
	}

	err := app.CreateTeam(context.TODO(), req)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}

func TestMysql_GetTeamByID_Success(t *testing.T) {
	app := teamRepoMysqlTest()

	resp, err := app.GetTeamByID(context.TODO(), 1)
	log.Printf("%+v", resp)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}