package redis

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/mwkosasih/soccer-service/repo/mysql"
	"github.com/stretchr/testify/assert"
)

func repoRedisTest() IRedis {
	// init env
	godotenv.Load("../../.env")

	// connect redis
	Connect()

	// connect mysql
	mysql.Connect()

	return NewRedis()
}

func TestRedis_Set_Success(t *testing.T) {
	app := repoRedisTest()

	key := "test"
	req := ""
	ttl := 10 * time.Second

	err := app.Set(context.TODO(), key, req, ttl)
	log.Printf("%+v", err)
	assert.NoErrorf(t, err, "This should have no error")
}

func TestRedis_Get_Success(t *testing.T) {
	app := repoRedisTest()

	key := "test"

	res, err := app.Get(context.TODO(), key)
	log.Printf("%+v", err)
	log.Printf("%+v", res)
	assert.NoErrorf(t, err, "This should have no error")
}
