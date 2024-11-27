package queue

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func parseTimestamp(ts int64, logs *log.Entry) int {
	format := time.Unix(ts, 0).Format("20060402")
	days, err := strconv.Atoi(format)
	if err != nil {
		logs.Errorf("时间转换失败: %v", err)
		return int(time.Now().Unix())
	}
	return days
}
