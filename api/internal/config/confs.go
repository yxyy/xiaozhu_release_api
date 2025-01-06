package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"path"
	"strings"
	"xiaozhu/utils"
)

var RootDir string

func Init() error {

	RootDir = utils.GetRunRootDir()

	// 设置配置文件名称
	viper.SetConfigName("conf")
	//	设置配置文件类型
	viper.SetConfigType("yaml")
	// 设置配置文件路径
	viper.AddConfigPath(RootDir)
	viper.AddConfigPath(path.Join(RootDir, "/etc"))

	// 加载环境变量配置（优先级高于YAML配置）
	err := godotenv.Load(RootDir + "/.env") // 加载 .env 文件（如果存在）
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// 启用自动读取环境变量并覆盖配置文件中的值
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // 支持点号转为下划线

	// 启用配置文件的热加载
	viper.WatchConfig()

	// 读取配置文件
	return viper.ReadInConfig()
}
