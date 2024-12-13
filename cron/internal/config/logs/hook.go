package logs

import (
	log "github.com/sirupsen/logrus"
	"xiaozhu/utils"
)

// ExtraDataHook 定义一个钩子，用来在每个日志记录中添加请求的字段
type ExtraDataHook struct {
	RequestID string
}

func (hook *ExtraDataHook) Levels() []log.Level {
	return log.AllLevels
}

func (hook *ExtraDataHook) Fire(entry *log.Entry) error {
	if entry.Data == nil {
		entry.Data = make(log.Fields)
	}

	if requestID, ok := entry.Context.Value("request_id").(string); ok {
		entry.Data["request_id"] = requestID
	} else {
		entry.Data["request_id"] = utils.Uuid()
	}

	// todo 第三方存储，如 kafka

	return nil
}
