package model

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/util"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type UserBase struct {
}

// 数据库User
type User struct {
	gorm.Model
	UserId   string
	UserName string
	TiZhi    int    // 体质
	MinJie   int    // 敏捷
	LingGen  string // 灵根
	Feature  string
	BaseInfo UserBase   `gorm:"embedded;embeddedPrefix:base_"`
	Dead     bool       // 死亡  true为死亡状态 默认为false
	StartAT  *time.Time //本次游戏开始时间
	DeadAt   *time.Time //死亡时间
	Life     int        //寿元
}

var lingGenList = []string{"金", "木", "水", "火", "土"}

// NewUser :create new user by name and id/*
func (u *User) NewUser(name string) {
	u.UserName = name
	setBaseInfo(u) // 设置基础属性值
	detail := BuildUserDetail(u)
	for _, i := range strings.Split(u.Feature, ",") {
		ExecCreateFeature(i, detail)
	}

}
func setBaseInfo(user *User) {
	user.TiZhi = util.RandomRange(20, 101) // 体质
	user.MinJie = util.RandomRange(1, 101) // 敏捷
	user.LingGen = GenLingGen()            // 灵根
	now := time.Now()
	user.StartAT = &now                  //开始时间
	user.Life = util.RandomRange(50, 81) //寿元
	user.Feature = "身强体壮"
}

// GenLingGen 生成灵根
func GenLingGen() string {
	percent := util.RandomDistribution(100, 5)
	var m = make(map[string]uint)
	for i := 0; i < 5; i++ {
		m[lingGenList[i]] = percent[i]
	}
	data, err := json.Marshal(&m)
	if err != nil {
		log.Println("map转换json出错， err = ", err)
		os.Exit(1)
	}
	return string(data)
}

func (u *User) Exist() bool {
	db.Where("user_id = ? and dead = false", u.UserId).First(u)
	if u.ID == 0 {
		return false
	}
	return true
}

func (u User) Create() {
	db.Create(&u)
}

func (u *User) UserInfo() {
	db.Where("user_id = ? and dead != true", u.UserId).First(u)
}

func ExistUserName(name string) bool {
	var user User
	db.Where("user_name = ? and dead != true", name).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

func (u *User) Save() {
	db.Save(&u)
}

// UserDetail User详情
type UserDetail struct {
	User *User
	//灵根map
	LingGenMap map[string]int
	//灵根描述（大于20且从大到小）
	LingGenDesc string
	//年龄
	Age int
	//剩余寿元
	LeftAge    int
	BattleInfo *UserBattleInfo
}

type UserBattleInfo struct {
	//血量
	HP int
	//灵力
	MP int
	//攻击力
	ATK int
	// 防御力
	DEF int
	// 速度
	SPD int
}

// 构建用户战斗信息
func BuildUserBattleInfo(user *User) *UserBattleInfo {
	var info UserBattleInfo
	//计算血量
	info.HP = util.IntReflect(user.TiZhi, 0, 100, 80, 130)
	//计算速度
	info.SPD = util.IntReflect(user.MinJie, 0, 100, 70, 100)
	info.MP = 0
	info.ATK = 10
	info.DEF = 0
	return &info

}

// BuildUserDetail 构建用户详情
func BuildUserDetail(user *User) *UserDetail {
	var detail UserDetail
	detail.User = user
	lingGenMap := make(map[string]int)
	err := json.Unmarshal([]byte(detail.User.LingGen), &lingGenMap)
	if err != nil {
		fmt.Println(err)
	}
	detail.LingGenMap = lingGenMap
	detail.LingGenDesc = GetLingGenDesc(lingGenMap)
	detail.Age = user.GetAge()
	detail.LeftAge = user.GetAgeLeft()
	detail.BattleInfo = BuildUserBattleInfo(user)
	for _, i := range strings.Split(user.Feature, ",") {
		ExecCalAttFeature(i, &detail)
	}
	return &detail
}

// GetLingGenDesc 获得20以上的灵根并从大到小排序
func GetLingGenDesc(lingGenMap map[string]int) string {
	var pairList PairList
	for k, v := range lingGenMap {
		if v >= 20 {
			pairList = append(pairList, Pair{k, v})
		}
	}
	sort.Sort(sort.Reverse(pairList))
	var lingGenStr string
	for _, e := range pairList {
		lingGenStr += e.Key
	}
	return lingGenStr
}

type Pair struct {
	Key   string
	Value int
}
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// GetAgeLeft 计算剩余寿元
func (u *User) GetAgeLeft() int {
	now := time.Now()
	return u.Life - int(now.Sub(*u.StartAT).Hours()/24)

}

// GetAge 获取当前年龄
func (u *User) GetAge() int {
	now := time.Now()
	return 20 + int(now.Sub(*u.StartAT).Hours()/24)

}

// GetDeadTime 获取寿元耗尽死亡时间
func (u *User) GetDeadTime() time.Time {
	d, _ := time.ParseDuration(string(u.Life*24) + "h")
	start := *u.StartAT
	t := start.Add(d)
	return t
}
