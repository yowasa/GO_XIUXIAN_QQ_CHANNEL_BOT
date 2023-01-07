package model

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
	"encoding/json"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	luajson "layeh.com/gopher-json"
	"log"
)

//type User struct {
//	Id       int64 `gorm:"primary_key" json:"id"`
//	Username string
//	Password string
//	Test     int64
//}

var db *gorm.DB
var config = cfg.GetConfig()

func init() {
	dsn := config.Mysql
	myDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myDb)
	myDb.AutoMigrate(
		User{},
	)
	db = myDb

}

// 执行
func execLua(user *UserDetail, luaPath string, method string) {
	u := *user
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	uJson := string(jsonBytes)

	// 创建一个lua解释器实例
	l := lua.NewState()
	luajson.Preload(l)
	// 在这个方法return后关闭lua解释器
	defer l.Close()
	err = l.DoFile(luaPath)
	if err != nil {
		log.Println(err)
	}

	// 执行具体的lua脚本
	err = l.CallByParam(lua.P{
		Fn:      l.GetGlobal(method), // 获取info函数引用
		NRet:    1,                   // 指定返回值数量
		Protect: true,                // 如果出现异常，是panic还是返回err
	}, lua.LString(uJson)) // 传递输入参数n=1
	if err != nil {
		panic(err)
	}
	// 获取返回结果
	ret := l.Get(-1)
	// 从堆栈中删除返回值
	l.Pop(1)

	err = json.Unmarshal([]byte(fmt.Sprint(ret)), &u)
	user = &u
}
