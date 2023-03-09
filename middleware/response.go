package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raven0520/wallet/app"
	"github.com/raven0520/wallet/util"
)

// Response Return Response
type Response struct {
	Code    util.ResponseCode `json:"code"`
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data"`
	TraceID string            `json:"trace_id"`
	Stack   interface{}       `json:"stack"`
}

// ResponseSignError Sign 错误
func ResponseSignError(context *gin.Context, err error) {
	TraceID := getTraceID(context)
	resp := &Response{Code: util.SignErrorCode, Status: "error", Message: err.Error(), TraceID: TraceID}
	context.JSON(200, resp)
	_ = context.AbortWithError(200, err)
}

// ResponseTokenError Token 错误
func ResponseTokenError(context *gin.Context, err error) {
	TraceID := getTraceID(context)
	resp := &Response{Code: util.TokenErrorCode, Status: "error", Message: err.Error(), TraceID: TraceID}
	context.JSON(200, resp)
	_ = context.AbortWithError(200, err)
}

// ResponseScopeError Scope 错误
func ResponseScopeError(context *gin.Context, err error) {
	TraceID := getTraceID(context)
	resp := &Response{Code: util.ScopeErrorCode, Status: "error", Message: err.Error(), TraceID: TraceID}
	context.JSON(200, resp)
	_ = context.AbortWithError(200, err)
}

// ResponseVerification 数据验证错误
func ResponseVerification(context *gin.Context, err error) {
	TraceID := getTraceID(context)
	resp := &Response{Code: util.Verification, Status: "warning", Message: err.Error(), TraceID: TraceID}
	context.JSON(200, resp)
	_ = context.AbortWithError(200, err)
}

// ResponseNotFound 未找到接口
func ResponseNotFound(context *gin.Context) {
	TraceID := getTraceID(context)
	resp := &Response{Code: util.NoFound, Status: "error", Message: util.ErrNotFound.Error(), TraceID: TraceID}
	context.JSON(200, resp)
	_ = context.AbortWithError(200, errors.New("NotFound"))
}

// ResponseSystemError 系统返回
func ResponseSystemError(context *gin.Context, err error) {
	TraceID := getTraceID(context)
	stack := getStack(context, err.Error())
	resp := &Response{Code: util.SystemError, Status: "error", Message: err.Error(), TraceID: TraceID, Stack: stack}
	context.JSON(200, resp)
	// _ = context.AbortWithError(200, err)
}

// ResponseError 错误返回
func ResponseError(context *gin.Context, message string, err error) {
	if err != nil {
		message = err.Error()
	}
	TraceID := getTraceID(context)
	stack := getStack(context, message)
	resp := &Response{Code: util.ErrorCode, Status: "error", Message: message, TraceID: TraceID, Stack: stack}
	context.JSON(200, resp)
	// _ = context.AbortWithError(200, err)
}

// ResponseSuccess 正确返回
func ResponseSuccess(context *gin.Context, data interface{}, message string) {
	TraceID := getTraceID(context)
	resp := &Response{Code: util.SuccessCode, Status: "success", Message: message, Data: data, TraceID: TraceID}
	context.JSON(200, resp)
}

// 获取链路ID
func getTraceID(context *gin.Context) string {
	_trace, _ := context.Get("trace")
	traceContext, _ := _trace.(*util.TContext)
	traceID := ""
	if traceContext != nil {
		traceID = traceContext.TraceID
	}
	return traceID
}

func getStack(context *gin.Context, err string) string {
	stack := ""
	if context.Query("debug_mode") == "debug" || app.GetEnv() == "Dev" {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err+"\n", "", -1)
	}
	return stack
}
