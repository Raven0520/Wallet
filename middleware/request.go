package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raven0520/wallet/util"
)

// RequestIn Build a request log structure
func RequestIn(context *gin.Context) {
	traceContext := util.NewTrace()
	if traceID := context.Request.Header.Get("com-header-rid"); traceID != "" {
		traceContext.TraceID = traceID
	}
	if spanID := context.Request.Header.Get("com-header-spanid"); spanID != "" {
		traceContext.SpanID = spanID
	}
	context.Set("startExecTime", time.Now())
	context.Set("trace", traceContext)

	bodyBytes, _ := io.ReadAll(context.Request.Body)
	context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back
	if len(bodyBytes) < 1024*1024 {
		context.Set("requestBody", string(bodyBytes))
	}
}

// RequestLog Request Log
func RequestLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		RequestIn(context)
		context.Next()
	}
}
