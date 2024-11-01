package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"xiaozhu/utils"

	log "github.com/sirupsen/logrus"
)

func InitLogs() error {

	// 设置格式
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               false,
		DisableColors:             true,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          true,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})

	logfile := RootDir + viper.GetString("logs.path")
	fmt.Println(logfile)
	err := utils.TidyDirectory(logfile)
	if err != nil {
		return err
	}

	fp, err := os.OpenFile(logfile+viper.GetString("logs.name"), os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	// defer fp.Close()
	// 	设置输出位置
	log.SetOutput(fp)
	//  设置日志等级
	log.SetLevel(log.InfoLevel)

	// 	TODO 分割日志

	return nil

}
