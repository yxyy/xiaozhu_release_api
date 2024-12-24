package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime"
	"strings"
)

const (
	SuccessCode = 20000
	ErrorCode   = 40000
	FailCode    = 40001
	ServerError = 50000

	SuccessMsg = "SUCCESS"
	FailMsg    = "Fail"
)

type Result struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	RequestId string `json:"request_id"`
}

type Response struct {
	Ctx    *gin.Context
	Result *Result
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
		Result: &Result{
			Code:      SuccessCode,
			Message:   SuccessMsg,
			Data:      nil,
			RequestId: ctx.GetString("request_id"),
		},
	}
}

func (r *Response) Success() {
	r.send(http.StatusOK)
}

func (r *Response) SuccessData(data any) {
	r.Result.Data = data
	r.send(http.StatusOK)
}

func (r *Response) Fail(message string) {
	r.Result.Code = FailCode
	r.Result.Message = message
	if r.Result.Message == "" {
		r.Result.Message = FailMsg
	}
	r.send(http.StatusBadRequest)
}

func (r *Response) Error(message error) {
	r.Result.Code = ErrorCode
	r.Result.Message = message.Error()
	r.send(http.StatusBadRequest)
}

func (r *Response) SetServerError(message string) {
	r.Result.Code = ServerError
	r.Result.Message = message
	r.send(http.StatusInternalServerError)
}

func (r *Response) SetCodeError(code int, message string) {
	r.Result.Code = code
	r.Result.Message = message
	if code == 401 {
		r.send(http.StatusUnauthorized)
	} else {
		r.send(http.StatusBadRequest)
	}

}

func (r *Response) SetResult(code int, message string, data any) {
	r.Result.Code = code
	r.Result.Message = message
	r.Result.Data = data
	r.send(code)
}

func (r *Response) send(httpStatus int) {
	r.Ctx.JSON(httpStatus, r.Result)
	// r.Ctx.Abort()

	logs := log.WithContext(r.Ctx)
	switch r.Result.Code {
	case ErrorCode, FailCode:
		StackTrace := r.logStackTrace()
		logs.WithField("StackTrace", StackTrace).Error(r.Result.Message)
	default:
		logs.Info(r.Result.Message)
	}
}

func (r *Response) OriginSuccess(message any) {
	r.originSend(http.StatusOK, message)
}

func (r *Response) OriginError(message any) {
	r.originSend(http.StatusBadRequest, message)
}

func (r *Response) SetOriginServerError(message any) {
	r.originSend(http.StatusInternalServerError, message)
}

func (r *Response) SetOriginResult(code int, data any) {
	r.originSend(code, data)
}

func (r *Response) originSend(httpStatus int, data any) {
	r.Ctx.JSON(httpStatus, data)

	logs := log.WithContext(r.Ctx)
	switch httpStatus {
	case ErrorCode, FailCode:
		StackTrace := r.logStackTrace()
		logs.WithField("StackTrace", StackTrace).Error(data)
	default:
		logs.Info(data)
	}
}

func (r *Response) logStackTrace() string {
	const maxStackDepth = 10
	var pcs [maxStackDepth]uintptr
	n := runtime.Callers(3, pcs[:]) // skip 3 to exclude logStackTrace, send, and Error itself

	var sb strings.Builder
	sb.WriteString("Stack trace:\n")
	for i := 0; i < n; i++ {
		pc := pcs[i]
		funcObj := runtime.FuncForPC(pc)
		file, line := funcObj.FileLine(pc)
		// 过滤只包含项目相关路径的调用栈信息
		if strings.Contains(file, "/internal/") {
			sb.WriteString(fmt.Sprintf("  at %s (%s:%d)\n", funcObj.Name(), file, line))
		}
	}

	return sb.String()
}
