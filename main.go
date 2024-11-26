package main

import (
	"tiktok/config"
	"tiktok/internal/repo/dao"
	"tiktok/internal/router"
)

func main() {
	config.InitConfig()
	dao.InitMysql()
	router.StartServer()
}
