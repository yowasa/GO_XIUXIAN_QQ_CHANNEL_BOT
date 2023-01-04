package com

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"context"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/openapi"
	"log"
)

type BotInfo struct {
	//大频道ID
	GuildID string
	//bot的api
	Api openapi.OpenAPI
	//bot的ctz
	Ctx context.Context
	//事件
	Event *dto.WSPayload
	//消息
	Data *dto.WSATMessageData
	//指令后追加的内容
	Content string
	//用户信息
	CurrentUser *model.User
	//AT的用户信息
	//atUser *model.User
}

// 回复子频道发送信息(引用回复)
func (bot *BotInfo) ReplayMsg(msg string) {
	bot.Api.PostMessage(bot.Ctx, bot.Data.ChannelID, &dto.MessageToCreate{MsgID: bot.Data.ID, Content: msg, MessageReference: &dto.MessageReference{MessageID: bot.Data.ID}})
}

// 回复子频道发送信息(不引用回复)
func (bot *BotInfo) ReplayMsgNotRef(msg string) {
	bot.Api.PostMessage(bot.Ctx, bot.Data.ChannelID, &dto.MessageToCreate{MsgID: bot.Data.ID, Content: msg})
}

// 主动向指定子频道发送信息
func (bot *BotInfo) SendMsg(channelID string, msg string) {
	bot.Api.PostMessage(bot.Ctx, channelID, &dto.MessageToCreate{Content: msg})
}

// 私信回复
func (bot *BotInfo) ReplyDirectMsg(msg string) {
	//创建私信会话
	directMsg, err := bot.Api.CreateDirectMessage(bot.Ctx, &dto.DirectMessageToCreate{
		SourceGuildID: bot.GuildID,
		RecipientID:   bot.Data.Author.ID,
	})
	if err != nil {
		log.Println("私信创建出错了，err = ", err)
	}
	bot.Api.PostDirectMessage(bot.Ctx, directMsg, &dto.MessageToCreate{MsgID: bot.Data.ID, Content: msg})
}

// 私信发送 腾讯侧会异步审核
func (bot *BotInfo) SendDirectMsg(userId string, msg string) {
	//创建私信会话
	directMsg, err := bot.Api.CreateDirectMessage(bot.Ctx, &dto.DirectMessageToCreate{
		SourceGuildID: bot.GuildID,
		RecipientID:   userId,
	})
	if err != nil {
		log.Println("私信创建出错了，err = ", err)
	}
	bot.Api.PostDirectMessage(bot.Ctx, directMsg, &dto.MessageToCreate{Content: msg})
}

var (
	//过滤at信息
	ATFilter = make(map[string]func(bot *BotInfo))
	//过滤私信信息
	DirectFilter = make(map[string]func(bot *BotInfo))
)

// 初始化 将指令与方法注册进去
func init() {
	ATFilter["/test"] = testFilter
}

func testFilter(botInfo *BotInfo) {
	botInfo.ReplayMsg("测试filter成功")
}
