package handler

import (
	"github.com/robfig/cron/v3"
)

func StartJobs() {

	c := cron.New()

	// c.AddFunc("*/1 * * * *", func() {
	// 	fmt.Println("Every hour on the half hour")
	// })

	// c.AddJob("*/1 * * * *", job.QMonitor)

	// ticker := time.NewTicker(10 * time.Second)
	// defer ticker.Stop()

	// for range ticker.C {
	// 	queue.QMonitor.Run()
	// }

	c.Start()

}
