package mysql

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/mwkosasih/soccer-service/entity"
	"github.com/mwkosasih/soccer-service/repo/redis"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func playerRepoMysqlTest() IPlayer {
	// init env
	godotenv.Load("../../.env")

	// connect redis
	redis.Connect()

	// connect mysql
	Connect()

	return NewPlayer()
}

func TestMysql_CreatePlayer_Success(t *testing.T) {
	app := playerRepoMysqlTest()

	req := entity.Player{
		TeamID:   1,
		Name:     "Wildan",
		Position: "Kiper",
		Height:   190,
		Weight:   30,
	}

	err := app.CreatePlayer(context.TODO(), req)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}

func TestMysql_GetPlayerByID_Success(t *testing.T) {
	app := playerRepoMysqlTest()

	resp, err := app.GetPlayerByID(context.TODO(), 1)
	log.Printf("%+v", resp)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}

func TestMysql_GetPlayerByTeamID_Success(t *testing.T) {
	app := playerRepoMysqlTest()

	resp, err := app.GetPlayerByTeamID(context.TODO(), 1)
	log.Printf("%+v", resp)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}