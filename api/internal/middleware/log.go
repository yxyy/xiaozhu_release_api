package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"xiaozhu/utils"
)

// 响应白名单
var paths = map[string]bool{
	"/system/v1/operation-log/list": true,
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // 将响应数据写入缓冲区
	return w.ResponseWriter.Write(b) // 将响应数据写入实际响应
}

func Log(c *gin.Context) {

	var err error
	var body []byte

	uuid := utils.Uuid()

	// 1
	withValue := context.WithValue(context.Background(), "request_id", uuid)
	c.Request = c.Request.WithContext(withValue)

	// 2
	c.Set("request_id", uuid)

	// log.AddHook(&utils.ExtraDataHook{RequestID: uuid})

	if c.Request.Method == "POST" {
		switch c.ContentType() {
		case "application/x-www-form-urlencoded":
			if err = c.Request.ParseForm(); err != nil {
				log.WithContext(c.Request.Context()).Error(err)
				return
			}
			body, err = json.Marshal(c.Request.Form)
			if err != nil {
				log.Error(err)
				return
			}

		case "application/json":
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				log.WithContext(c.Request.Context()).Error(err)
				return
			}
			// 重写回去
			c.Request.Body = io.NopCloser(bytes.NewReader(body))

		}
	}

	// 使用自定义 ResponseWriter 替换gin的响应接口
	writer := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = writer

	c.Next()

	responseBody := ""
	if !paths[c.Request.URL.Path] {
		responseBody = writer.body.String()
	}

	go func() {
		logger := log.WithFields(log.Fields{
			"request_id":   uuid,
			"ip":           c.ClientIP(),
			"method":       c.Request.Method,
			"url":          c.Request.URL.Path,
			"Access-Token": c.Request.Header.Get("Access-Token"),
			"response":     responseBody,
		})

		if c.Request.Method == "POST" {
			logger = logger.WithFields(log.Fields{
				"ContentType": c.ContentType(),
				"body":        string(body),
			})
		}
		logger.WithContext(c.Request.Context()).Info("请求日志")
	}()

	// go func() {
	// 	path := c.Request.URL.Path
	// 	moduleIndex := strings.Index(path[1:], "/")
	// 	module := path[1 : moduleIndex+1]
	// 	typeIndex := strings.LastIndex(path, "/")
	// 	businessPath := path[typeIndex:]
	//
	// 	logs := user.SysUserLog{
	// 		LogType:   getLogType(businessPath),
	// 		UserId:    c.GetInt("userId"),
	// 		Account:   c.GetString("account"),
	// 		Module:    module,
	// 		Ip:        c.ClientIP(),
	// 		Path:      path,
	// 		UserAgent: c.Request.UserAgent(),
	// 		Request:   string(body),
	// 		Response:  responseBody,
	// 		Status:    writer.Status(),
	// 		RequestId: c.GetString("request_id"),
	// 	}
	//
	// 	if err = logs.Create(); err != nil {
	// 		log.Error(err)
	// 	}
	// }()

}
