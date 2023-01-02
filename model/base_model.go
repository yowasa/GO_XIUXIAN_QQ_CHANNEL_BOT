package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int64 `gorm:"primary_key" json:"id"`
	Username string
	Password string
}

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	db.AutoMigrate(
		User{},
	)
}
