package main

import (
	"tiktok/common/logger"
	"tiktok/config"
	"tiktok/internal/repo/dao"
	"tiktok/internal/router"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	dao.InitMysql()
	router.StartServer()
}
