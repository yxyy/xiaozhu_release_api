package handler

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func StartJobs() {

	c := cron.New()

	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Every hour on the half hour")
	})

	c.Start()

}
