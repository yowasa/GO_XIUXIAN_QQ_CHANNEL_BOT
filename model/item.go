package model

import "gorm.io/gorm"

type UserItem struct {
	gorm.Model
	UserId   uint   //归属人员
	ItemId   int    //道具id
	ItemName string //道具名称
	Attr     string //道具属性
	Num      int    //道具数量
}

func (ui *UserItem) Save() {
	db.Create(ui)
}

func SearchItem(userId uint, name string) UserItem {
	var ui UserItem
	db.Where("user_id = ? and item_name = ?", userId, name).First(&ui)
	return ui
}
