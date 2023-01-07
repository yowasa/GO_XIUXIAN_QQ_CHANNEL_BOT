package model

import (
	"GO_XIUXIAN_QQ_CHANNEL_BOT/cfg"
)

// ExecCreateFeature 创建角色时特征效果
func ExecCreateFeature(featureName string, user *UserDetail) {
	luaPath := cfg.PathFeatureLua[featureName]
	execLua(user, luaPath, "inCreate")
}

// ExecCalAttFeature 常驻加成
func ExecCalAttFeature(featureName string, user *UserDetail) {
	luaPath := cfg.PathFeatureLua[featureName]
	execLua(user, luaPath, "inCalAttr")
}
