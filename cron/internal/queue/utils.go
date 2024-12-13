package queue

import (
	"fmt"
	"strconv"
	"time"
)

func parseTimestamp(ts int64) (int, int64, error) {
	if ts <= 0 {
		ts = time.Now().UnixMilli()
	}
	format := time.Unix(ts/1000, 0).Format("20060402")
	days, err := strconv.Atoi(format)
	if err != nil {
		return 0, ts, fmt.Errorf("时间转换失败: %v", err)
	}
	return days, ts, nil
}
