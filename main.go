package main

import (
	"github.com/mwkosasih/soccer-service/repo/mysql"
	"github.com/mwkosasih/soccer-service/repo/redis"
	"github.com/mwkosasih/soccer-service/transport/grpc/config"
	"github.com/mwkosasih/soccer-service/util"
)

func main() {
	// init env
	util.Env("./")

	// connect redis
	redis.Connect()

	// connect mysql
	mysql.Connect()

	// run migration
	mysql.RunMigrate()

	// serve grpc
	config.Connect()
}
