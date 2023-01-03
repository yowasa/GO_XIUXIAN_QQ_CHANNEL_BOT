package model

// 数据库User
type User struct {
	Id       int64 `gorm:"primary_key" json:"id"`
	UserId   string
	UserName string
	Con      int64
	Agi      int64
	BaseInfo string
}

// NewUser :create new user by name and id/*
func NewUser(name string, userId string) *User {
	var user User = User{
		UserName: name,
		UserId:   userId,
	}
	// todo 基本属性值
	return &user
}

func (u User) Exist() bool {
	db.Where("UserId = ?", u.UserId).First(&u)
	if u.Id == 0 {
		return false
	}
	return true
}

func (u User) Create() {
	db.Create(&u)
}

func (u User) ExistName(name string) bool {
	db.Where("UserName = ?", name).First(&u)
	if u.Id == 0 {
		return false
	}
	return true
}

// todo
func getLingGen() string {

	return ""
}

func (u User) Save() {
	db.Save(&u)
}

// User详情
type UserDetail struct {
	User User
}
