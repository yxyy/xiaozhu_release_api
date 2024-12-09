package handler

import (
	"github.com/robfig/cron/v3"
	"xiaozhu/internal/job"
)

func StartJobs() {

	c := cron.New()

	// c.AddFunc("*/1 * * * *", func() {
	// 	fmt.Println("Every hour on the half hour")
	// })

	c.AddJob("*/1 * * * *", job.DefaultMonitor)

	c.Start()

}
