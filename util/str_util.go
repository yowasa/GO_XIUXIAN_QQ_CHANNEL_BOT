package util

import "github.com/tencent-connect/botgo/dto"

// BuildEmbed 获取 Embed
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
