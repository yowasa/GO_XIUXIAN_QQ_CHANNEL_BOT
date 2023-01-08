package model

var Channels map[string]Channel

type Channel struct {
	//Id              int64  `gorm:"primary_key" json:"id"`
	ChannelId       string // 子频道 id
	ChannelName     string // 子频道名称/身份组名称
	RoleId          string // 对应可访问用户组 id
	ParentChannelId string //最顶层channel Id  或者说取一个类型
}
