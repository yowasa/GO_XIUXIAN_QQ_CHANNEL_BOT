package com

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/util"
	"fmt"
)

// CreateUserFilter 创建新用户
func CreateUserFilter(bot *BotInfo) {
	user := bot.CurrentUser
	if user.ExistName(botInfo.Content) {
		bot.ReplayMsg("该角色名已存在，请更换角色名")
		return
	}
	user.NewUser(bot.Content)
	user.Save()
	bot.ReplayMsg(fmt.Sprintf("拥有%s灵根的%s已经进入修仙界", user.LingGen, user.UserName))
}

// personalInfoFilter 展示个人信息
func personalInfoFilter(botInfo *BotInfo) {
	var user = botInfo.CurrentUser
	user.UserInfo()
	var info = []string{
		user.UserName,
		string(user.TiZhi),
		string(user.MinJie),
		user.LingGen,
	}
	// todo
	botInfo.ReplayEmbedMsg(util.BuildEmbed("个人信息", botInfo.Data.Author.Avatar, info))
}
