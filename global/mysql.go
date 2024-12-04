package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Db *gorm.DB

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		Conf.Mysql.Username,
		Conf.Mysql.Password,
		Conf.Mysql.Host,
		Conf.Mysql.Port,
		Conf.Mysql.Database,
		Conf.Mysql.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("Open mysql error: ", err.Error())
	}

	db.Logger.LogMode(logger.Info)

	Db = db
}
