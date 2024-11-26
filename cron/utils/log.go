package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

func InitLogs() error {

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

	// log.AddHook(&ExtraDataHook{})  在日志中间件加

	// 分割日志
	go cutting()

	return nil

}

func setOut() error {

	logPath := path.Join(RootDir, path.Clean(viper.GetString("logs.path")))
	if err := TidyDirectory(logPath); err != nil {
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

func cutting() {

	mod := viper.GetString("logs.mod")
	filename := viper.GetString("logs.name")
	logPath := path.Join(RootDir, path.Clean(viper.GetString("logs.path"))+"/")

	// 根据模式设置时间间隔
	var duration time.Duration
	switch mod {
	case "minute":
		duration = time.Minute
	case "hour":
		duration = time.Hour
	case "days":
		duration = time.Hour * 24
	default:
		fmt.Println("日志切割模式未启用,无需切割日志")
	}
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	fmt.Println("日志分割开始准备完成....")
	for range ticker.C {
		fmt.Println("旧日志名称准备中....")
		format := ""
		switch mod {
		case "minute":
			format = time.Now().Add(-duration).Format("20060102_1504")
		case "hour":
			format = time.Now().Add(-duration).Format("20060102_15")
		case "days":
			format = time.Now().Add(-duration).Format("20060102")
		default:
			fmt.Println("无需切割日志")
		}
		fmt.Println("旧日志名称准备完成....")

		oldName := path.Join(logPath, filename+".log")
		newName := path.Join(logPath, fmt.Sprintf("%s_%s.log", filename, format))

		fmt.Println("检查旧日志....")
		// 检查旧日志文件是否存在
		if _, err := os.Stat(oldName); os.IsNotExist(err) {
			fmt.Println("日志文件不存在，跳过切割:", oldName)
			continue
		}

		// 先关闭旧文件
		CloseLogs()

		fmt.Println("开始重命名....")
		err := os.Rename(oldName, newName)
		if err != nil {
			fmt.Println("重命名失败....")
			log.Error(err)
			return
		}

		fmt.Println("开始重命名完成....")

		// 	重新打开文件
		if err = setOut(); err != nil {
			fmt.Println("重新打开文件失败....", err)
		}

	}

}

// ExtraDataHook 定义一个钩子，用来在每个日志记录中添加请求的字段
type ExtraDataHook struct {
	RequestID string
}

func (hook *ExtraDataHook) Levels() []log.Level {
	return log.AllLevels
}

func (hook *ExtraDataHook) Fire(entry *log.Entry) error {

	if entry.Data == nil {
		entry.Data = make(log.Fields)
	}

	// 在每条日志记录中加入 request_id 字段
	entry.Data["request_id"] = hook.RequestID
	return nil
}
