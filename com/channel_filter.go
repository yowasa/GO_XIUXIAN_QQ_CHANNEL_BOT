package com

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"github.com/spf13/cast"
	"github.com/tencent-connect/botgo/dto"
	"math"
	"time"
)

// moveFilter 移动
func moveFilter(bot *BotInfo) {
	user := bot.CurrentUser
	if !user.CheckFree() {
		bot.ReplyMsg(user.GetStatusMsg())
		return
	}
	var destination = bot.Content
	if len(destination) == 0 {
		bot.ReplyMsg("请在指令后输入要移动到的目的地")
		return
	}
	channel := model.Channels[destination]
	if channel == (model.Channel{}) {
		bot.ReplyMsg("所输入的目的地不存在，请输入正确的目的地")
		return
	}
	// todo 是否有权限可以去

	info := model.BuildUserBattleInfo(user)
	var needTime = moveTime(info.SPD, user.Location, destination)
	// 状态更改
	user.SetStatusBusy(1, "移动")
	bot.ReplyMsg("你开始移动中,大约需要时间" + cast.ToString(needTime) + "分钟")
	go moveToChannel(needTime, bot, channel.RoleId, user, destination)

	endTime := time.Now().Add(time.Minute * time.Duration(needTime))
	msg := " 你已移动到:" + destination
	model.NewEvent(*user, channel.ChannelId, bot.Data.ID, msg, endTime)

}

func moveToChannel(needTime int, bot *BotInfo, roleId string, user *model.User, destination string) {
	time.Sleep(time.Minute * time.Duration(needTime))
	// 移动到新子频道
	bot.Api.MemberAddRole(bot.Ctx, bot.GuildID, dto.RoleID(roleId), user.UserId, nil)
	// 从旧子频道中移除
	channel := model.Channels[user.Location]
	bot.Api.MemberDeleteRole(bot.Ctx, bot.GuildID, dto.RoleID(channel.RoleId), user.UserId, nil)
	user.Location = destination
	user.Save()
	user.SetStatusFree()
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

func moveTime(speed int, A string, B string) int {
	locationA := cfg.MapLocation[A]
	locationB := cfg.MapLocation[B]
	var distance = math.Sqrt(math.Pow(math.Abs(locationA.X-locationB.X), 2) + math.Pow(math.Abs(locationA.Y-locationB.Y), 2))
	return int(math.Ceil(distance / float64(speed)))
}
