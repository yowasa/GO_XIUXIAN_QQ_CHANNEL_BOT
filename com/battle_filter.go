package com

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/battle"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/util"
)

// battleFilter 展示个人信息
func battleFilter(botInfo *BotInfo) {
	user := botInfo.CurrentUser
	atUserId := botInfo.AtUserList[0]
	var B = model.User{
		UserId: atUserId,
	}
	if !B.Exist() {
		botInfo.ReplayMsg("对手尚未开始修仙")
	}
	var msg = battle.Battle(user, &B)
	botInfo.ReplyDirectEmbedMsg(util.BuildEmbed("战斗概况", botInfo.Data.Author.Avatar, msg))

}
