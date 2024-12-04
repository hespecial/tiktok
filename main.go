package main

import (
	"tiktok/global"
	"tiktok/internal/router"
)

func main() {
	global.InitConfig()
	global.InitLogger()
	global.InitMysql()
	global.InitRedis()
	router.StartServer()
}
