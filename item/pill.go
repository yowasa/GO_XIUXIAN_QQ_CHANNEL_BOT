package item

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/model"
	"encoding/json"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	luajson "layeh.com/gopher-json"
	"log"
)

/*丹药处理*/

type PillContent struct {
	User   *model.UserDetail
	Item   *model.UserItem
	Result bool
	Msg    string
}

// 执行
func execLua(content *PillContent, luaPath string, method string) {
	c := *content
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	cJson := string(jsonBytes)

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
	}, lua.LString(cJson)) // 传递输入参数n=1
	if err != nil {
		return
	}
	// 获取返回结果
	ret := l.Get(-1)
	// 从堆栈中删除返回值
	l.Pop(1)
	//var newUser UserDetail
	err = json.Unmarshal([]byte(fmt.Sprint(ret)), &c)
	content = &c
}
