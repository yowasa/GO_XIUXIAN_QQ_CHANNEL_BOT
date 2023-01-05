package battle

import "GO_XIUXIAN_QQ_CHANNEL_BOT/model"

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

// 获取下一个回合是谁动
func (SP *SpeedProgress) NextTurn() int {
	//aNext:=float64(SP.MaxIndex-SP.AIndex/float64(SP.ASpeed))
	return 0
}
