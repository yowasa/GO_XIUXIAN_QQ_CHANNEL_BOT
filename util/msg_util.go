package util

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"regexp"
	"strings"
)

// BuildEmbed 组装 Embed
func BuildEmbed(title string, picUrl string, msgList []string) *dto.Embed {
	var fieldList []*dto.EmbedField
	for _, msg := range msgList {
		fieldList = append(fieldList, &dto.EmbedField{Name: msg})
	}
	return &dto.Embed{
		Title: title,
		Thumbnail: dto.MessageEmbedThumbnail{
			URL: picUrl,
		},
		Fields: fieldList,
	}
}

// GetAtList 获取at的用户列表
func GetAtList(str string) []string {
	res := message.ParseCommand(str) //去掉@结构和清除前后空格
	cmd := res.Cmd                   ///对于像 /私信天气 城市名 指令，cmd 为 私信天气
	start := strings.Index(str, cmd)
	subStr := str[start : len(str)-1]
	var atRE = regexp.MustCompile(`<@!\d+>`)
	atMsg := atRE.FindAllString(subStr, -1)
	if len(atMsg) == 0 {
		return nil
	}
	var result []string
	for _, msg := range atMsg {
		msg = strings.ReplaceAll(msg, "<@!", "")
		msg = strings.ReplaceAll(msg, ">", "")
		result = append(result, msg)
	}
	return result
}

// GetAtList1 获取除了at机器人之后at的用户
func GetAtList1(users []*dto.User) []string {
	var atList []string
	for i := 1; i < len(users); i++ {
		atList = append(atList, users[i].ID)
	}
	return atList
}

// GetFirstAt 获取除了at机器人之后第一个at的用户
func GetFirstAt(str string) string {
	atList := GetAtList(str)
	if atList == nil {
		return ""
	}
	return atList[0]
}
