package com

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"GO_XIUXIAN_QQ_CHANNEL_BOT/util"
	"fmt"
)

var (
	// ATFilter 过滤at信息
	ATFilter = make(map[string]func(bot *BotInfo))
	// DirectFilter 过滤私信信息
	DirectFilter = make(map[string]func(bot *BotInfo))
)

// 初始化 将指令与方法注册进去
func init() {
	ATFilter["个人信息"] = personalInfoFilter
	DirectFilter["/个人信息"] = personalInfoFilter
	ATFilter["开始修炼"] = getExpByPractice
	ATFilter["停止修炼"] = stopGetExpByPractice
	ATFilter["升级"] = levelUp
	ATFilter["突破"] = stageUp
	ATFilter["使用"] = useItem

	//====== done ========
	ATFilter["获取丹药"] = getPill
	ATFilter["/test"] = testFilter
	ATFilter["移动"] = moveFilter
	ATFilter["战斗"] = battleFilter
	// 创建者需要具有相应管理员权限
	ATFilter["创建子频道"] = createChannelFilter
	ATFilter["创建身份组"] = createRoleFilter

	DirectFilter["/战斗"] = battleFilter
	DirectFilter["/移动"] = moveFilter

	DirectFilter["/创建子频道"] = createChannelFilter
}

func testFilter(botInfo *BotInfo) {
	pinzhi := []string{"下品", "中品", "上品", "极品"}
	p := pinzhi[util.RandomN(4)]
	i := cfg.ItemNameMap["培元丹"]
	AddPill(botInfo.CurrentUser.ID, i, p)
	botInfo.ReplyMsg(fmt.Sprintf("获取%s丹药成功", i.Name))
}

func getPill(botInfo *BotInfo) {
	pinzhi := []string{"下品", "中品", "上品", "极品"}
	p := pinzhi[util.RandomN(4)]
	i := cfg.ItemNameMap["培元丹"]
	AddPill(botInfo.CurrentUser.ID, i, p)
	botInfo.ReplyMsg(fmt.Sprintf("获取%s丹药成功", i.Name))
}

// AddPill 获得丹药 quality：=[]string{"下品", "中品", "上品", "极品"}选一个
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
