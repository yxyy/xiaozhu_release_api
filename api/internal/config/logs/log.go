package logs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"xiaozhu/internal/config"
	"xiaozhu/utils"
)

func Init() error {

	// 设置格式
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		QuoteEmptyFields: true,
	})

	if err := setOut(); err != nil {
		return err
	}

	//  设置日志等级
	log.SetLevel(log.InfoLevel)

	// log.Hooks = make(log.LevelHooks)

	log.AddHook(&ExtraDataHook{})

	// 分割日志
	go cutter()

	return nil

}

func setOut() error {

	logPath := path.Join(config.RootDir, path.Clean(viper.GetString("logs.path")))
	if err := utils.TidyDirectory(logPath); err != nil {
		return err
	}

	logfile := path.Join(logPath, viper.GetString("logs.name")+".log")
	fp, err := os.OpenFile(logfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	// defer fp.Close()

	// 	设置输出位置
	log.SetOutput(fp)

	return nil
}

func CloseLogs() {
	// if file, ok := log.Out.(*os.File); ok {
	// 	_ = file.Close()
	// }

	if file, ok := log.StandardLogger().Out.(*os.File); ok {
		err := file.Close()
		if err != nil {
			log.Error(err)
			fmt.Println("关闭日志文件失败", err)
		}
	}

}
