package jobs

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func InitJobs(c *cron.Cron) {

	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Every hour on the half hour")
	})

}
