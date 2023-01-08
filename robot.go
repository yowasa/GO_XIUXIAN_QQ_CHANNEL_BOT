package main

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/com"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/util"
	"context"
	"github.com/robfig/cron"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"log"
	"os"
	"strings"
	"time"
)

var config = cfg.GetConfig()
var api openapi.OpenAPI
var ctx context.Context

func directMessageEventHandler(event *dto.WSPayload, data *dto.WSDirectMessageData) error {
	res := message.ParseCommand(data.Content) //去掉@结构和清除前后空格
	log.Println("cmd = " + res.Cmd + " content = " + res.Content)
	cmd := res.Cmd         ///对于像 /私信天气 城市名 指令，cmd 为 私信天气
	content := res.Content //content 为 城市名
	var user model.User
	user.UserId = data.Author.ID
	myData := dto.Message(*data)
	myBot := com.BotInfo{
		GuildID:     config.GuildId,
		Api:         api,
		Ctx:         ctx,
		Event:       event,
		Data:        &myData,
		Content:     content,
		CurrentUser: &user,
	}
	if !user.Exist() {
		myBot.ReplayMsg("请先创建角色再执行指令！")
		return nil
	}
	if com.DirectFilter[cmd] != nil {
		com.DirectFilter[cmd](&myBot)
	}

	return nil
}

// atMessageEventHandler 处理 @机器人 的消息
func atMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	res := message.ParseCommand(data.Content) //去掉@结构和清除前后空格

	log.Println("cmd = " + res.Cmd + " content = " + res.Content)
	cmd := res.Cmd         ///对于像 /私信天气 城市名 指令，cmd 为 私信天气
	content := res.Content //content 为 城市名
	//atList := util.GetAtList(data.Content)
	atList := util.GetAtList1(data.Mentions)
	var user model.User
	user.UserId = data.Author.ID
	myData := dto.Message(*data)
	myBot := com.BotInfo{
		GuildID:     config.GuildId,
		Api:         api,
		Ctx:         ctx,
		Event:       event,
		Data:        &myData,
		Content:     content,
		CurrentUser: &user,
		AtUserList:  atList,
	}
	if !user.Exist() {
		if strings.EqualFold(cmd, "开始修仙") {
			com.CreateUserFilter(&myBot)
			return nil
		}
		myBot.ReplayMsg("请先创建角色再执行指令！")
		return nil
	}
	if com.ATFilter[cmd] != nil {
		com.ATFilter[cmd](&myBot)
	}

	return nil
}

func main() {

	//第二步：生成token，用于校验机器人的身份信息
	token := token.BotToken(config.AppID, config.Token)
	//第三步：获取操作机器人的API对象
	api = botgo.NewOpenAPI(token).WithTimeout(3 * time.Second)
	//获取context
	ctx = context.Background()
	//第四步：获取websocket
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalln("websocket错误， err = ", err)
		os.Exit(1)
	}

	registerMsgPush()
	// 初始化所有的子频道及用户组到常量map中
	initChannel()
	var atMessage event.ATMessageEventHandler = atMessageEventHandler
	var directMessage event.DirectMessageEventHandler = directMessageEventHandler

	intent := websocket.RegisterHandlers(atMessage, directMessage) // 注册socket消息处理
	botgo.NewSessionManager().Start(ws, token, &intent)            // 启动socket监听
}

// registerMsgPush 注册定时器
func registerMsgPush() {
	var activeMsgPush = func() {
		channelId := config.TestChannelId
		if channelId != "" {
			//MsgID 为空字符串表示主动消息
			api.PostMessage(ctx, channelId, &dto.MessageToCreate{MsgID: "", Content: "当前天气是：晴天"})
		}
	}
	timer := cron.New()
	//cron表达式由6部分组成，从左到右分别表示 秒 分 时 日 月 星期
	//*表示任意值  ？表示不确定值，只能用于星期和日
	//这里表示每天15:53分发送消息
	timer.AddFunc("0 53 15 * * ?", activeMsgPush)
	timer.Start()
}

func initChannel() {

	channels, err := api.Channels(ctx, config.GuildId)
	if err != nil {
		log.Fatalln("get channel list error， err = ", err)
		os.Exit(1)
	}
	roles, err := api.Roles(ctx, config.GuildId)
	if err != nil {
		log.Fatalln("get role list error， err = ", err)
		os.Exit(1)
	}
	model.Roles = make(map[string]model.Role, len(channels))
	for i := 0; i < len(roles.Roles); i++ {
		var name = roles.Roles[i].Name
		var role = model.Role{
			RoleId:   string(roles.Roles[i].ID),
			RoleName: roles.Roles[i].Name,
		}
		model.Roles[name] = role
	}
	model.Channels = make(map[string]model.Channel, len(channels))
	for i := 0; i < len(channels); i++ {
		var name = channels[i].Name
		model.Channels[name] = model.Channel{
			ChannelId:       channels[i].ID,
			ChannelName:     channels[i].Name,
			ParentChannelId: channels[i].ParentID,
			RoleId:          model.Roles[name].RoleId,
		}
	}

}
