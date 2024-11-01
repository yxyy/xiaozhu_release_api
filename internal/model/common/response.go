package common

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	RequestId string      `json:"request_id"`
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

func (r *Response) SuccessData(data interface{}) {
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
	r.send(http.StatusBadRequest)
}

func (r *Response) SetResult(code int, message string, data interface{}) {
	r.Result.Code = code
	r.Result.Message = message
	r.Result.Data = data
	r.send(code)
}

func (r *Response) send(httpStatus int) {
	r.Ctx.JSON(httpStatus, r.Result)
	r.Ctx.Abort()

	logs := log.WithField("request_id", r.Result.RequestId)
	switch r.Result.Code {
	case ErrorCode:
		logs.Error(r.Result.Message)
	case FailCode:
		logs.Fatal(r.Result.Message)
	default:
		logs.Info(r.Result.Message)
	}
}
