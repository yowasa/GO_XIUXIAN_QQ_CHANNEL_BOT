package model

type User struct {
	Id       int64 `gorm:"primary_key" json:"id"`
	UserId   string
	UserName string
	BaseInfo string
}

func (u User) Exist() bool {
	db.Where("UserId = ?", u.UserId).First(&u)
	if u.Id == 0 {
		return false
	}
	return true

}
