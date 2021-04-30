package server

import (
	"context"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mwkosasih/soccer-service/repo/mysql"
	"github.com/mwkosasih/soccer-service/repo/redis"
	pb "github.com/mwkosasih/soccer-service/transport/grpc/proto/soccer"
	"github.com/stretchr/testify/assert"
)

func serverTest() *Server {
	// init env
	godotenv.Load("../../.env")

	// connect redis
	redis.Connect()

	// connect mysql
	mysql.Connect()

	return NewServer()
}

func TestServer_CreateTeam_Success(t *testing.T) {
	app := serverTest()
	req := pb.CreateTeamRequest{
		Name: "Manchester United",
	}
	res, err := app.CreateTeam(context.TODO(), &req)
	log.Printf("%+v", res)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}

func TestServer_CreatePlayer_Success(t *testing.T) {
	app := serverTest()
	req := pb.CreatePlayerRequest{
		TeamId:   int32(1),
		Name:     "Manchester United",
		Position: "Kiper",
		Height:   int32(170),
		Weight:   int32(70),
	}
	res, err := app.CreatePlayer(context.TODO(), &req)
	log.Printf("%+v", res)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}

func TestServer_GetPlayers_Success(t *testing.T) {
	app := serverTest()
	req := pb.GetPlayerRequest{
		Id: 1,
	}
	res, err := app.GetPlayer(context.TODO(), &req)
	log.Printf("%+v", res)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}
