package com

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/util"
	"fmt"
	"github.com/spf13/cast"
)

// CreateUserFilter 创建新用户
func CreateUserFilter(bot *BotInfo) {
	user := bot.CurrentUser
	if len(bot.Content) == 0 {
		bot.ReplayMsg("请在指令后加上创建人物的名称")
		return
	}
	if model.ExistUserName(bot.Content) {
		bot.ReplayMsg("该角色名已存在，请更换角色名")
		return
	}
	user.NewUser(bot.Content)
	user.Save()
	detail := model.BuildUserDetail(user)
	bot.ReplayMsg(fmt.Sprintf("拥有%s灵根,特性%s的%s已经进入修仙界", detail.LingGenDesc, user.Feature, user.UserName))
}

// personalInfoFilter 展示个人信息
func personalInfoFilter(botInfo *BotInfo) {
	var user = botInfo.CurrentUser
	detail := model.BuildUserDetail(user)
	var info = []string{
		//user.UserName,
		"体质: " + cast.ToString(user.TiZhi) + "\t" + "敏捷: " + cast.ToString(user.MinJie),
		"灵根: " + detail.LingGenDesc,
		"年龄: " + cast.ToString(detail.Age) + "\t" + "寿元: " + cast.ToString(detail.LeftAge),
		"位置: " + cast.ToString(user.Location),
		fmt.Sprintf("特性: %s", detail.User.Feature),
		fmt.Sprintf("HP: %d\tMP: %d\tSPD: %d\t", detail.BattleInfo.HP, detail.BattleInfo.MP, detail.BattleInfo.SPD),
		fmt.Sprintf("ATK:%d\tDEF:%d\t", detail.BattleInfo.ATK, detail.BattleInfo.DEF),
	}
	botInfo.ReplyDirectEmbedMsg(util.BuildEmbed(user.UserName, botInfo.Data.Author.Avatar, info))
}
