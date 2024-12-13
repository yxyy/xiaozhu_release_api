package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"net/http"
	_ "net/http/pprof"
	"xiaozhu/internal/config"
	"xiaozhu/internal/router"
	utilsqueue "xiaozhu/utils/queue"
)

const defaultPort = "80"

func main() {
	ServerRun()
}

func ServerRun() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	Init()
	r := router.InitRouter()
	port := viper.GetString("port")
	if port == "" {
		port = defaultPort
	}
	if err := r.Run(":" + port); r != nil {
		log.Fatal("服务启动失败： %w", err)
	}
}

func Init() {
	if err := config.InitConf(); err != nil {
		log.Fatalln("配置初始失败：", err)
	}

	// 初始化日志
	if err := config.InitLogs(); err != nil {
		log.Fatalln("日志初始失败：", err)
	}
	// defer utils.CloseLogs()

	if err := config.InitMysql(); err != nil {
		log.Fatalln("MYSQL初始失败：", err)
	}

	if err := config.InitRedis(); err != nil {
		log.Fatalln("redis初始失败：", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置文件: %s 发生变化,Op %d: \n", e.Name, e.Op)
	})

	// 注册队列连接器
	redis := &utilsqueue.Redis{Conn: config.RedisDB00}
	utilsqueue.RegisterCoupler(redis)
}
