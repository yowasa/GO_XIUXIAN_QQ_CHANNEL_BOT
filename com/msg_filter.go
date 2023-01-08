package com

var (
	// ATFilter 过滤at信息
	ATFilter = make(map[string]func(bot *BotInfo))
	// DirectFilter 过滤私信信息
	DirectFilter = make(map[string]func(bot *BotInfo))
)

// 初始化 将指令与方法注册进去
func init() {
	ATFilter["私信"] = privateFilter
	ATFilter["/test"] = testFilter
	ATFilter["个人信息"] = personalInfoFilter
	ATFilter["移动"] = moveFilter
	ATFilter["战斗"] = battleFilter
	// 创建者需要具有相应管理员权限
	ATFilter["创建子频道"] = createChannelFilter
	ATFilter["创建身份组"] = createRoleFilter

	DirectFilter["/个人信息"] = personalInfoFilter
	DirectFilter["/战斗"] = battleFilter
	DirectFilter["/移动"] = moveFilter

	DirectFilter["/创建子频道"] = createChannelFilter
}

func testFilter(botInfo *BotInfo) {
	botInfo.ReplayMsg("测试filter成功")
}

func privateFilter(botInfo *BotInfo) {
	botInfo.ReplyDirectMsg("你好，欢迎来到修仙世界")
}
