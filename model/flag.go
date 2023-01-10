package model

import "gorm.io/gorm"

type UserFlag struct {
	gorm.Model
	UserId   int    //人员
	FlagCode string //标记code
	Value    int    //标记值
	Attr     string //标记额外参数
}
