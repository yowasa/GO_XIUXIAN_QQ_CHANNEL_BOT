package cfg

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path"
)

// Config 全局配置
type Config struct {
	AppID         uint64 `yaml:"appid"`         //机器人的appid
	Token         string `yaml:"token"`         //机器人的token
	GuildId       string `yaml:"guildId"`       //频道id
	TestChannelId string `yaml:"testChannelId"` //测试自频道id
	Mysql         string `yaml:"mysql"`         //mysql数据库链接dsn
}

// 全局配置
var config Config

// Level 境界配置
type Level struct {
	Radio int `json:"radio""`
	Exp   int `json:"exp"`
}

// UserLevel 境界映射
var UserLevel map[string]Level

// PathFeatureLua 特征名-lua脚本位置映射
var PathFeatureLua map[string]string

// PathMentalLua 心法名-lua脚本位置映射
var PathMentalLua map[string]string

// PathSpecialLua 特技名-lua脚本位置映射
var PathSpecialLua map[string]string

func GetConfig() *Config {
	return &config
}

func init() {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Println("读取配置文件出错， err = ", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Println("解析配置文件出错， err = ", err)
		os.Exit(1)
	}
	levelPath := "./cfg/file/json/level.json"
	content, err = os.ReadFile(levelPath)
	err = json.Unmarshal(content, &UserLevel)
	if err != nil {
		log.Println("解析配置文件出错， err = ", err)
		os.Exit(1)
	}

	featurePath := "./cfg/file/lua/feature_skill"
	PathFeatureLua, err = GetAllFile(featurePath)
	if err != nil {
		fmt.Println("read dir fail:", err)
	}
	mentalPath := "./cfg/file/lua/mental_skill"
	PathMentalLua, err = GetAllFile(mentalPath)
	if err != nil {
		fmt.Println("read dir fail:", err)
	}
	specialPath := "./cfg/file/lua/special_skill"
	PathSpecialLua, err = GetAllFile(specialPath)
	if err != nil {
		fmt.Println("read dir fail:", err)
	}

}

// GetAllFile 获取路径下所有文件名和文件路径的映射
func GetAllFile(pathname string) (map[string]string, error) {
	var dic = make(map[string]string)
	rd, err := os.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return dic, err
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := pathname + "/" + fi.Name()
			dic[GetFilePrefix(fullName)] = fullName
		}
	}
	return dic, nil
}

// GetFilePrefix 获得文件不包含类型后缀的纯名称
func GetFilePrefix(filepath string) string {
	filenameall := path.Base(filepath)
	filesuffix := path.Ext(filepath)
	fileprefix := filenameall[0 : len(filenameall)-len(filesuffix)]
	return fileprefix
}
