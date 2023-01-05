package battle

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"github.com/spf13/cast"
)

// MAX_BATTLE_COUNT 最大战斗次数
const MAX_BATTLE_COUNT = 20

// Content 战斗ctx
type Content struct {
	//第一个战斗人员
	A *model.UserDetail
	//第二个战斗人员
	B *model.UserDetail
	//速度进度条
	SPED *SpeedProgress
	//

}

// SpeedProgress 进度条
type SpeedProgress struct {
	//进度框长度
	MaxIndex int
	//A的速度
	ASpeed int
	//B的速度
	BSpeed int
	//A当前位置
	AIndex int
	//B当前位置
	BIndex int
}

func (content *Content) initSpeedProgress() {
	content.SPED = &SpeedProgress{
		100,
		content.A.BattleInfo.SPD,
		content.B.BattleInfo.SPD,
		0,
		0,
	}

}

// 获取下一个回合是谁动
func (SP *SpeedProgress) NextTurn() bool {
	aTime := float64(SP.MaxIndex-SP.AIndex) / float64(SP.ASpeed)
	bTime := float64(SP.MaxIndex-SP.BIndex) / float64(SP.BSpeed)
	// 下回合状态
	var result = aTime < bTime
	if result {
		SP.BIndex += int(float64(SP.BSpeed) * aTime)
		SP.AIndex = 0
	} else {
		SP.AIndex += int(float64(SP.ASpeed) * bTime)
		SP.BIndex = 0
	}
	return result
}

func buildContent(a *model.User, b *model.User) *Content {
	var content Content
	content.A = model.BuildUserDetail(a)
	content.B = model.BuildUserDetail(b)
	content.A.BattleInfo = model.BuildUserBattleInfo(a)
	content.B.BattleInfo = model.BuildUserBattleInfo(b)
	// speed info
	content.initSpeedProgress()
	return &content
}

func Battle(a *model.User, b *model.User) []string {
	var content = buildContent(a, b)
	var ABattleInfo = content.A.BattleInfo
	var BBattleInfo = content.B.BattleInfo

	// start battle
	// which one is 先动手
	// 战斗停止条件 ： 战斗次数，或者某个认的HP为0，或者到达某个条件
	var msg []string
	var battleCount = 0
	for battleCount <= MAX_BATTLE_COUNT {
		battleCount++
		// 为true a先出手; false b先出手
		if content.SPED.NextTurn() {
			// todo
			msg = append(msg, battleProcess(content.A, content.B))
		} else {
			msg = append(msg, battleProcess(content.B, content.A))
		}
		if ABattleInfo.HP > 0 && BBattleInfo.HP > 0 {
			continue
		}
		if ABattleInfo.HP <= 0 && BBattleInfo.HP <= 0 {
			msg = append(msg, a.UserName+"与"+b.UserName+"同归于尽")
		} else if ABattleInfo.HP <= 0 {
			msg = append(msg, a.UserName+"死啦")
		} else if BBattleInfo.HP <= 0 {
			msg = append(msg, b.UserName+"死啦")
		}
		break
	}
	msg = append(msg, b.UserName+"战斗回合已至上限，战斗结束")
	return msg
}

func battleProcess(A *model.UserDetail, B *model.UserDetail) string {
	var damage = A.BattleInfo.ATK - B.BattleInfo.DEF
	// todo
	B.BattleInfo.HP -= damage
	//inMsg := *msg
	//inMsg = append(inMsg, "")
	//msg = &inMsg
	return A.User.UserName + "使用了普通攻击，" + B.User.UserName + "受到了" + cast.ToString(damage) + "伤害," + B.User.UserName + "剩余HP为" + cast.ToString(B.BattleInfo.HP)
}
