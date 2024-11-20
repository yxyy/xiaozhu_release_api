package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var RootDir string

func InitConf() error {
	var err error
	RootDir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 设置配置文件名称
	viper.SetConfigName("conf")
	//	设置配置文件类型
	viper.SetConfigType("yaml")
	// 设置配置文件路径
	viper.AddConfigPath(RootDir)
	viper.AddConfigPath(RootDir + "/etc")

	// 启用配置文件的热加载
	viper.WatchConfig()

	// 读取配置文件
	return viper.ReadInConfig()
}
