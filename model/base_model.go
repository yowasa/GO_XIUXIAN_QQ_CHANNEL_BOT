package model

import (
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

func GetDB() *gorm.DB {
	return db
}

func init() {
	dsn := "root:123456@tcp(81.69.241.164:3306)/xiuxian_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	db.AutoMigrate(
		User{},
	)
}
