package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {

	fmt.Println("start xiaozhu corn ...")

	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Every hour on the half hour")
	})

	c.Start()

	select {}
}
