package com

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/battle"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/util"
)

// battleFilter 展示个人信息
func battleFilter(botInfo *BotInfo) {
	var A model.User
	var B model.User
	A.UserId = "3649196412093161908"
	B.UserId = "12396374757621072114"
	A.Exist()
	B.Exist()
	var msg = battle.Battle(&A, &B)
	botInfo.ReplyDirectEmbedMsg(util.BuildEmbed("战斗", botInfo.Data.Author.Avatar, msg))

}
