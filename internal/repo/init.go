package repo

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"tiktok/config"
)

var Db *gorm.DB

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("Open mysql error: ", err.Error())
	}

	db.Logger.LogMode(logger.Info)

	Db = db
}

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Conf.Redis.Host, config.Conf.Redis.Port),
		DB:   config.Conf.Redis.Database,
	})
	if err := Redis.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Init redis error: ", err.Error())
	}
}
