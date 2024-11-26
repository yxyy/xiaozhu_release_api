package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"log"
	"xiaozhu/cron/internal/jobs"
	"xiaozhu/cron/utils"
)

func main() {
	Init()

	c := cron.New()
	jobs.InitJobs(c)

	c.Start()

	fmt.Println("start xiaozhu corn ...")

	select {}
}

func Init() {

	if err := utils.InitConf(); err != nil {
		log.Fatalln("配置初始失败：", err)
	}

	// 初始化日志
	if err := utils.InitLogs(); err != nil {
		log.Fatalln("日志初始失败：", err)
	}
	// defer utils.CloseLogs()

	if err := utils.InitMysql(); err != nil {
		log.Fatalln("MYSQL初始失败：", err)
	}

	if err := utils.InitRedis(); err != nil {
		log.Fatalln("redis初始失败：", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置文件: %s 发生变化,Op %d: \n", e.Name, e.Op)
	})
}
