package model

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//type User struct {
//	Id       int64 `gorm:"primary_key" json:"id"`
//	Username string
//	Password string
//	Test     int64
//}

var db *gorm.DB
var config = cfg.GetConfig()

func GetDB() *gorm.DB {
	return db
}

func init() {
	dsn := config.Mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	db.AutoMigrate(
		User{},
	)
}
