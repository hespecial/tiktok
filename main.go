package main

import (
	"tiktok/common/logger"
	"tiktok/config"
	"tiktok/internal/repo"
	"tiktok/internal/router"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	repo.InitMysql()
	repo.InitRedis()
	router.StartServer()
}
