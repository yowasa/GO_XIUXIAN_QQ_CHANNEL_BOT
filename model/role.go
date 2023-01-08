package model

var Roles map[string]Role

type Role struct {
	Id       int64  `gorm:"primary_key" json:"id"`
	RoleId   string // 身份组id
	RoleName string // 身份组id
}
