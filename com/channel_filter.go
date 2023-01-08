package com

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"github.com/tencent-connect/botgo/dto"
)

// moveFilter 移动
func moveFilter(bot *BotInfo) {
	var destination = bot.Content
	if len(destination) == 0 {
		bot.ReplayMsg("请在指令后输入要移动到的目的地")
		return
	}
	channel := model.Channels[destination]
	if channel == (model.Channel{}) {
		bot.ReplayMsg("所输入的目的地不存在，请输入正确的目的地")
		return
	}
	user := bot.CurrentUser
	// todo 是否有权限可以去

	moveToChannel(bot, channel.RoleId, user, destination)
	// todo 速度时间计算
	bot.ReplayMsg("你已移动到:" + destination)
}

func moveToChannel(bot *BotInfo, roleId string, user *model.User, destination string) {
	// 移动到新子频道
	bot.Api.MemberAddRole(bot.Ctx, bot.GuildID, dto.RoleID(roleId), user.UserId, nil)
	// 从旧子频道中移除
	channel := model.Channels[user.Location]
	bot.Api.MemberDeleteRole(bot.Ctx, bot.GuildID, dto.RoleID(channel.RoleId), user.UserId, nil)
	user.Location = destination
}

func createChannelFilter(bot *BotInfo) {
	var channel = dto.ChannelValueObject{
		Name:     bot.Content,
		ParentID: "0",
		Position: 10,
	}
	bot.Api.PostChannel(bot.Ctx, bot.GuildID, &channel)
}

func createRoleFilter(bot *BotInfo) {
	var role = dto.Role{
		Name: bot.Content,
	}
	bot.Api.PostRole(bot.Ctx, bot.GuildID, &role)
}
