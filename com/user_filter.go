package com

import "GO_XIUXIAN_QQ_CHANNEL_BOT/util"

// createUserFilter 创建用户角色
func createUserFilter(botInfo *BotInfo) {
	var user = botInfo.CurrentUser
	if user.Exist() {
		botInfo.ReplayMsg("您的角色已存在，请勿重复创建")
		return
	}
	if user.ExistName(botInfo.Content) {
		botInfo.ReplayMsg("该角色名已存在，请更换角色名")
	} else {
		user.NewUser(botInfo.Content)
		user.Create()
		botInfo.ReplayMsg("角色: " + botInfo.Content + " 创建成功")
	}
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
