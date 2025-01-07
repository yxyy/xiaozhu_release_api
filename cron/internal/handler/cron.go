package handler

import (
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"xiaozhu/internal/job"
)

const (
	orderJobName = "orderJob"
)

var (
	name    string
	rootCmd = &cobra.Command{
		Use:   "cron", // 根命令的名称
		Short: "这是我的命令行工具",
	}
	c           = cron.New()
	EntryIDName = make(map[string]cron.EntryID)
)

func StartJobs() {
	orderJobId, err := c.AddJob("*/1 * * * *", &job.OrderJob{})
	if err != nil {
		log.Errorf("启动定时任务失败：%s", err)
	}
	EntryIDName[orderJobName] = orderJobId
	cmdStart()

	c.Start()

}

func cmdStart() {
	rootCmd.AddCommand(runCmd())
	rootCmd.AddCommand(listCmd())

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func listCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "查看所有任务",
		Run: func(cmd *cobra.Command, args []string) {
			for k, v := range EntryIDName {
				fmt.Printf("%-2d-------------------- %-20s\n", v, k)
			}
			os.Exit(0)
		},
	}
}

func runCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "run",
		Short: "执行任务",
		Run: func(cmd *cobra.Command, args []string) {
			jobName := cmd.Flag("name").Value.String()
			entry := c.Entry(EntryIDName[jobName])
			if entry.Job != nil {
				entry.Job.Run()
			}
			os.Exit(0)
		},
	}

	// 添加命令行参数
	command.Flags().StringVarP(&name, "name", "n", "", "job name")

	return command
}
