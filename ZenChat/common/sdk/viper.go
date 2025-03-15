package sdk

import (
	"github.com/fsnotify/fsnotify"
	"github.com/mazezen/itools"
	"github.com/spf13/viper"
)

var cf *ChatConfig

type ChatConfig struct {
	HttpPort string              `json:"http_port"`
	Db       itools.DbOption     `json:"db"`
	Redis    itools.RedisOption  `json:"redis"`
	Logger   itools.LoggerOption `json:"logger"`
	Jwt      Jwt                 `json:"jwt"`
}

type Jwt struct {
	Expire int    `json:"expire"`
	Secret string `json:"secret"`
}

func init() {
	cf = &ChatConfig{}
}

func GetConf() *ChatConfig {
	return cf
}

func ParseConfig(fileName string) {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	v.WatchConfig()
	rc := func() {
		if err := v.Unmarshal(cf); err != nil {

			panic(err)
		}
	}
	v.OnConfigChange(func(in fsnotify.Event) {
		rc()
	})
	rc()
}
