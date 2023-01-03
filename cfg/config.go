package cfg

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	AppID         uint64 `yaml:"appid"`         //机器人的appid
	Token         string `yaml:"token"`         //机器人的token
	GuildId       uint64 `yaml:"guildId"`       //频道id
	TestChannelId string `yaml:"testChannelId"` //测试自频道id
	Mysql         string `yaml:"mysql"`         //mysql数据库链接dsn
}

var config Config

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
	log.Println(config)
}
