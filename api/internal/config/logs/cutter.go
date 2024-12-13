package logs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
	"xiaozhu/internal/config"
)

func cutter() {

	mod := viper.GetString("logs.mod")
	filename := viper.GetString("logs.name")
	logPath := path.Join(config.RootDir, path.Clean(viper.GetString("logs.path"))+"/")

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
