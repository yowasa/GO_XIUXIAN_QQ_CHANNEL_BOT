package main

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/com"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"context"
	"encoding/json"
	"github.com/robfig/cron"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// WeatherResp 定义了返回天气数据的结构
type WeatherResp struct {
	Success    string `json:"success"` //标识请求是否成功，0表示成功，1表示失败
	ResultData Result `json:"result"`  //请求成功时，获取的数据
	Msg        string `json:"msg"`     //请求失败时，失败的原因
}

// Result 定义了具体天气数据结构
type Result struct {
	Days            string `json:"days"`             //日期，例如2022-03-01
	Week            string `json:"week"`             //星期几
	CityNm          string `json:"citynm"`           //城市名
	Temperature     string `json:"temperature"`      //当日温度区间
	TemperatureCurr string `json:"temperature_curr"` //当前温度
	Humidity        string `json:"humidity"`         //湿度
	Weather         string `json:"weather"`          //天气情况
	Wind            string `json:"wind"`             //风向
	Winp            string `json:"winp"`             //风力
	TempHigh        string `json:"temp_high"`        //最高温度
	TempLow         string `json:"temp_low"`         //最低温度
	WeatherIcon     string `json:"weather_icon"`     //气象图标
}

var config = cfg.GetConfig()
var api openapi.OpenAPI
var ctx context.Context

// 定义常量
const (
	CmdDirectChatMsg        = "/私信天气"
	CmdNowWeather           = "/当前天气"
	CmdTestAtRpNormal       = "#tn"
	CmdTestAtRpEmbed        = "#te"
	CmdTestAtRpDirectNormal = "#tdn"
	CmdTestAtRpDirectEmbed  = "#tde"
	CmdTestDirectRpNormal   = "#dtn"
	CmdTestDirectRpEmbed    = "#dte"
	CmdTestUserExist        = "#exist"
)

func directMessageEventHandler(event *dto.WSPayload, data *dto.WSDirectMessageData) error {
	if strings.HasSuffix(data.Content, "hello") {
		directMsg, err := api.CreateDirectMessage(ctx, &dto.DirectMessageToCreate{
			SourceGuildID: data.GuildID,
			RecipientID:   data.Author.ID,
		})
		if err != nil {
			log.Println("私信创建出错了，err = ", err)
		}
		//发送私信消息
		api.PostDirectMessage(ctx, directMsg, &dto.MessageToCreate{MsgID: data.ID, Content: "你好"})
	}
	return nil
}

// return msg todo
func sendMsg(data *dto.WSATMessageData, msg string) {
	api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: msg})

}

// atMessageEventHandler 处理 @机器人 的消息
func atMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	res := message.ParseCommand(data.Content) //去掉@结构和清除前后空格
	log.Println("cmd = " + res.Cmd + " content = " + res.Content)
	cmd := res.Cmd         ///对于像 /私信天气 城市名 指令，cmd 为 私信天气
	content := res.Content //content 为 城市名
	switch cmd {
	case CmdTestAtRpNormal:
		api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "测试at普通回复成功"})
	case CmdTestAtRpEmbed:
		api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{Embed: createTestEmbed(data.Author.Avatar)})
	case CmdTestAtRpDirectNormal:
		//创建私信会话
		directMsg, err := api.CreateDirectMessage(ctx, &dto.DirectMessageToCreate{
			SourceGuildID: data.GuildID,
			RecipientID:   data.Author.ID,
		})
		if err != nil {
			log.Println("私信创建出错了，err = ", err)
		}
		api.PostDirectMessage(ctx, directMsg, &dto.MessageToCreate{MsgID: data.ID, Content: "测试at私信回复成功"})
	case CmdTestAtRpDirectEmbed:
		//创建私信会话
		directMsg, err := api.CreateDirectMessage(ctx, &dto.DirectMessageToCreate{
			SourceGuildID: data.GuildID,
			RecipientID:   data.Author.ID,
		})
		if err != nil {
			log.Println("私信创建出错了 ，err = ", err)
		}
		api.PostDirectMessage(ctx, directMsg, &dto.MessageToCreate{Embed: createTestEmbed(data.Author.Avatar)})

	case CmdTestUserExist:
		//var db = model.GetDB()
		var user model.User
		user.UserId = data.Author.ID
		if !user.Exist() {
			if user.ExistName(content) {
				sendMsg(data, "用户名已存在")
			} else {
				user.NewUser(content)
				user.Create()
				sendMsg(data, "角色: "+content+" 已创建")
			}
		} else {
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Content: "用户存在"})
		}

	case CmdNowWeather: //获取当前天气 指令是 /天气 城市名
		webData := getWeatherByCity(content)
		if webData != nil {
			//MsgID 表示这条消息的触发来源，如果为空字符串表示主动消息
			//Ark 传入数据时表示发送的消息是Ark
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{MsgID: data.ID, Ark: createArkForTemplate23(webData)})
		}
	case CmdDirectChatMsg: //私信天气消息到用户
		webData := getWeatherByCity(content)
		if webData != nil {
			//创建私信会话
			directMsg, err := api.CreateDirectMessage(ctx, &dto.DirectMessageToCreate{
				SourceGuildID: data.GuildID,
				RecipientID:   data.Author.ID,
			})
			if err != nil {
				log.Println("私信创建出错了，err = ", err)
			}
			//发送私信消息
			//Embed 传入数据时表示发送的是 Embed
			api.PostDirectMessage(ctx, directMsg, &dto.MessageToCreate{Embed: createEmbed(webData)})
		}
	default:
		var user model.User
		user.UserId = data.Author.ID
		myBot := com.BotInfo{
			GuildID:     config.GuildId,
			Api:         api,
			Ctx:         ctx,
			Event:       event,
			Data:        data,
			Content:     content,
			CurrentUser: &user,
		}
		if !user.Exist() {
			myBot.ReplayMsg("请先创建角色再执行指令！")
		}
		if com.ATFilter[cmd] != nil {
			com.ATFilter[cmd](&myBot)
		}
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

	var atMessage event.ATMessageEventHandler = atMessageEventHandler
	var directMessage event.DirectMessageEventHandler = directMessageEventHandler

	intent := websocket.RegisterHandlers(atMessage, directMessage) // 注册socket消息处理
	botgo.NewSessionManager().Start(ws, token, &intent)            // 启动socket监听
}

// 获取对应城市的天气数据
func getWeatherByCity(cityName string) *WeatherResp {
	url := "http://api.k780.com/?app=weather.today&cityNm=" + cityName + "&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("天气预报接口请求异常, err = ", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("天气预报接口数据异常, err = ", err)
		return nil
	}
	var weatherData WeatherResp
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Fatalln("解析数据异常 err = ", err, body)
		return nil
	}
	if weatherData.Success != "1" {
		log.Fatalln("返回数据问题 err = ", weatherData.Msg)
		return nil
	}
	return &weatherData
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

func createTestEmbed(url string) *dto.Embed {
	return &dto.Embed{
		Title: "测试成功",
		Thumbnail: dto.MessageEmbedThumbnail{
			URL: url,
		},
		Fields: []*dto.EmbedField{
			{
				Name: "测试列1",
			},
			{
				Name: "测试列2",
			},
		},
	}
}

// 获取 Embed
func createEmbed(weather *WeatherResp) *dto.Embed {
	return &dto.Embed{
		Title: weather.ResultData.CityNm + " " + weather.ResultData.Weather,
		Thumbnail: dto.MessageEmbedThumbnail{
			URL: weather.ResultData.WeatherIcon,
		},
		Fields: []*dto.EmbedField{
			{
				Name: weather.ResultData.Days + " " + weather.ResultData.Week,
			},
			{
				Name: "当日温度区间：" + weather.ResultData.Temperature,
			},
			{
				Name: "当前温度：" + weather.ResultData.TemperatureCurr,
			},
			{
				Name: "最高温度：" + weather.ResultData.TempHigh,
			},
			{
				Name: "最低温度：" + weather.ResultData.TempLow,
			},
			{
				Name: "当前湿度：" + weather.ResultData.Humidity,
			},
		},
	}
}

// 创建23号的Ark
func createArkForTemplate23(weather *WeatherResp) *dto.Ark {
	return &dto.Ark{
		TemplateID: 23,
		KV:         createArkKvArray(weather),
	}
}

// 创建Ark需要的ArkKV数组
func createArkKvArray(weather *WeatherResp) []*dto.ArkKV {
	akvArray := make([]*dto.ArkKV, 3)
	akvArray[0] = &dto.ArkKV{
		Key:   "#DESC#",
		Value: "描述",
	}
	akvArray[1] = &dto.ArkKV{
		Key:   "#PROMPT#",
		Value: "#PROMPT#",
	}
	akvArray[2] = &dto.ArkKV{
		Key: "#LIST#",
		Obj: createArkObjArray(weather),
	}
	return akvArray
}

// 创建ArkKV需要的ArkObj数组
func createArkObjArray(weather *WeatherResp) []*dto.ArkObj {
	objectArray := []*dto.ArkObj{
		{
			[]*dto.ArkObjKV{
				{
					Key:   "desc",
					Value: weather.ResultData.CityNm + " " + weather.ResultData.Weather + " " + weather.ResultData.Days + " " + weather.ResultData.Week,
				},
			},
		},
		{
			[]*dto.ArkObjKV{
				{
					Key:   "desc",
					Value: "当日温度区间：" + weather.ResultData.Temperature,
				},
			},
		},
		{
			[]*dto.ArkObjKV{
				{
					Key:   "desc",
					Value: "当前温度：" + weather.ResultData.TemperatureCurr,
				},
			},
		},
		{
			[]*dto.ArkObjKV{
				{
					Key:   "desc",
					Value: "当前湿度：" + weather.ResultData.Humidity,
				},
			},
		},
	}
	return objectArray
}
