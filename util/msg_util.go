package util

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
	"github.com/tencent-connect/botgo/dto"
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

// GetAtList 获取除了at机器人其他的at的用户
func GetAtList(users []*dto.User) []string {
	var atList []string
	for i := 0; i < len(users); i++ {
		if !strings.EqualFold(cfg.GetConfig().BotId, users[i].ID) {
			atList = append(atList, users[i].ID)
		}
	}
	return atList
}
