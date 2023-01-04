package com

var (
	// ATFilter 过滤at信息
	ATFilter = make(map[string]func(bot *BotInfo))
	// DirectFilter 过滤私信信息
	DirectFilter = make(map[string]func(bot *BotInfo))
)

// 初始化 将指令与方法注册进去
func init() {
	ATFilter["/test"] = testFilter
	ATFilter["/创建角色"] = createUserFilter
	ATFilter["/个人信息"] = personalInfoFilter
}

func testFilter(botInfo *BotInfo) {
	botInfo.ReplayMsg("测试filter成功")
}
