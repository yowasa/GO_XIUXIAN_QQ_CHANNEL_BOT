package com

import (
	"fmt"
)

// CreateUserFilter 创建新用户
func CreateUserFilter(bot *BotInfo) {
	user := bot.CurrentUser
	user.NewUser(bot.Content)
	user.Save()
	bot.ReplayMsg(fmt.Sprintf("拥有%s灵根的%s已经进入修仙界", user.LingGen, user.UserName))
}
