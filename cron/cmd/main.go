package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	_ "net/http/pprof"
	"xiaozhu/internal/config"
	"xiaozhu/internal/config/cache"
	"xiaozhu/internal/config/logs"
	"xiaozhu/internal/config/mysql"
	"xiaozhu/internal/handler"
)

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6061", nil))
	// }()
	ServerInit()
	handler.StartJobs()
	handler.StartQueue()
	fmt.Println("start xiaozhu corn ...")
	select {}
}

func ServerInit() {

	if err := config.InitConf(); err != nil {
		log.Fatalln("配置初始失败：", err)
	}

	// 初始化日志
	if err := logs.Init(); err != nil {
		log.Fatalln("日志初始失败：", err)
	}
	// defer utils.CloseLogs()

	if err := mysql.Init(); err != nil {
		log.Fatalln("MYSQL初始失败：", err)
	}

	if err := cache.InitRedis(); err != nil {
		log.Fatalln("redis初始失败：", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置文件: %s 发生变化,Op %d: \n", e.Name, e.Op)
	})

}
