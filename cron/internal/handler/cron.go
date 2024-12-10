package handler

import (
	"github.com/robfig/cron/v3"
	"time"
	"xiaozhu/internal/job"
)

func StartJobs() {

	c := cron.New()

	// c.AddFunc("*/1 * * * *", func() {
	// 	fmt.Println("Every hour on the half hour")
	// })

	// c.AddJob("*/1 * * * *", job.DefaultMonitor)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		job.DefaultMonitor.Run()
	}

	c.Start()

}
