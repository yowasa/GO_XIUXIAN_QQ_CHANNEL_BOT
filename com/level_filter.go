package com

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
	"fmt"
	"time"
)

/**
修炼模块
常规的获取经验和突破方式
*/

// 修炼获取经验
func getExpByPractice(botInfo *BotInfo) {
	u := botInfo.CurrentUser
	if !u.CheckFree() {
		botInfo.ReplyMsg(u.GetStatusMsg())
		return
	}
	u.BaseInfo.Status = 1
	now := time.Now()
	u.BaseInfo.StatusStartTime = &now
	botInfo.ReplyMsg("你开始进行修炼")
	u.Save()
}

// 停止修炼
func stopGetExpByPractice(botInfo *BotInfo) {
	u := botInfo.CurrentUser
	if u.BaseInfo.Status != 1 {
		botInfo.ReplyMsg("你没有开始修炼无需结束")
		return
	}
	u.BaseInfo.Status = 0
	now := time.Now()
	startTime := *u.BaseInfo.StatusStartTime
	//每小时等于半个月 一天能获取3点经验
	getExp := int(now.Sub(startTime).Hours() * 15)
	u.BaseInfo.NowExp += getExp
	u.Save()
	botInfo.ReplyMsg(fmt.Sprintf("经过修炼%s", u.GetExpMsg()))
}

// 升级
func levelUp(botInfo *BotInfo) {
	u := botInfo.CurrentUser
	stage := cfg.UserLevel[int(u.Stage/100)-1]
	needExp := stage.Exp
	if u.BaseInfo.NowExp < needExp {
		botInfo.ReplyMsg(fmt.Sprintf("你当前的修炼程度还不足，%s", u.GetExpMsg()))
		return
	}
	if u.Level >= len(stage.Level) {
		botInfo.ReplyMsg(fmt.Sprintf("你已经到达了%s的极限，请使用突破来突破当前境界", stage.Name))
		return
	}
	u.BaseInfo.NowExp = 0
	u.Level += 1
	u.Save()
	level := stage.Level[u.Level-1]
	botInfo.ReplyMsg(fmt.Sprintf("经过刻苦的修炼你终于达到了%s%s", stage.Name, level.Name))

}

// 突破
func stageUp(botInfo *BotInfo) {
	u := botInfo.CurrentUser
	stage := cfg.UserLevel[int(u.Stage/100)-1]
	needExp := stage.Exp
	if u.BaseInfo.NowExp < needExp {
		botInfo.ReplyMsg(fmt.Sprintf("你当前的修炼程度还不足，%s", u.GetExpMsg()))
		return
	}
	if u.Level < len(stage.Level) {
		botInfo.ReplyMsg(fmt.Sprintf("你还没有到达%s的最高等级，请使用升级来提升自己", stage.Name))
		return
	}
	u.BaseInfo.NowExp = 0
	u.Level = 1
	u.Stage += 100
	u.Save()
	stage = cfg.UserLevel[int(u.Stage/100)-1]
	level := stage.Level[u.Level-1]
	botInfo.ReplyMsg(fmt.Sprintf("经过刻苦的修炼你终于突破到了%s%s", stage.Name, level.Name))

}

//顿悟

//灌体

//心魔
