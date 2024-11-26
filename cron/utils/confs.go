package utils

import (
	"github.com/spf13/viper"
	"path"
)

var RootDir string

func InitConf() error {

	RootDir = GetRunRootDir()

	// 设置配置文件名称
	viper.SetConfigName("conf")
	//	设置配置文件类型
	viper.SetConfigType("yaml")
	// 设置配置文件路径
	viper.AddConfigPath(RootDir)
	viper.AddConfigPath(path.Join(RootDir, "/etc"))

	// 启用配置文件的热加载
	viper.WatchConfig()

	// 读取配置文件
	return viper.ReadInConfig()
}
