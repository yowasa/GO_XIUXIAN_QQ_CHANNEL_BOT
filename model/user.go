package model

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/util"
	"encoding/json"
	"log"
	"os"
)

// 数据库User
type User struct {
	Id       int64 `gorm:"primary_key" json:"id"`
	UserId   string
	UserName string
	TiZhi    int    // 体质
	MinJie   int    // 敏捷
	LingGen  string // 灵根
	BaseInfo string
	Dead     bool // 死亡  true为死亡状态
}

var lingGenList = []string{"金", "木", "水", "火", "土"}

// NewUser :create new user by name and id/*
func (u *User) NewUser(name string) {
	u.UserName = name
	setBaseInfo(u) // 设置基础属性值
}
func setBaseInfo(user *User) {
	user.TiZhi = util.RandomRange(20, 100) // 体质
	user.MinJie = util.RandomRange(0, 100) // 敏捷
	user.LingGen = getLingGen()            // 灵根
	user.Dead = true
}

func getLingGen() string {
	percent := util.RandomDistribution(100, 5)
	var m = make(map[string]uint)
	for i := 0; i < 5; i++ {
		m[lingGenList[i]] = percent[i]
	}
	data, err := json.Marshal(&m)
	if err != nil {
		log.Println("map转换json出错， err = ", err)
		os.Exit(1)
	}
	return string(data)
}

func (u User) Exist() bool {
	db.Where("user_id = ? and dead != 0", u.UserId).First(&u)
	if u.Id == 0 {
		return false
	}
	return true
}

func (u User) Create() {
	db.Create(&u)
}

func (u *User) UserInfo() {
	db.Where("user_id = ? and dead != 0", u.UserId).First(&u)
}

func (u User) ExistName(name string) bool {
	db.Where("user_name = ?", name).First(&u)
	if u.Id == 0 {
		return false
	}
	return true
}

func (u User) Save() {
	db.Save(&u)
}

// User详情
type UserDetail struct {
	User User
}
