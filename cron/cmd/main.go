package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"net/http"
	_ "net/http/pprof"
	"xiaozhu/internal/handler"
	"xiaozhu/utils"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6061", nil))
	}()
	ServerInit()
	// handler.StartJobs()
	handler.StartQueue()
	fmt.Println("start xiaozhu corn ...")
	select {}
}

func ServerInit() {

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
