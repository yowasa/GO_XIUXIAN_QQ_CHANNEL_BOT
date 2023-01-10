package item

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
)

/*丹药处理*/

// AddPill 获得丹药
func AddPill(userId uint, item cfg.Item, quality string) {
	fullName := quality + item.Name
	ui := model.SearchItem(userId, fullName)
	ui = model.UserItem{
		UserId:   userId,
		ItemId:   item.Id,
		ItemName: fullName,
		Num:      ui.Num + 1,
		Attr:     quality,
	}
	ui.Save()
}

func UsePill(u *model.User, fullName string) (bool, string) {
	ui := model.SearchItem(u.ID, fullName)
	if ui.Num <= 0 {
		return false, "道具数量不足"
	}
	return true, "todo@yowasa"
}
