package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/raven0520/wallet/util"
)

// GetGinTraceContext 从 gin 的 Context 中获取数据
func GetGinTraceContext(context *gin.Context) *util.TContext {
	// 防御
	if context == nil {
		return util.NewTrace()
	}
	traceContext, exists := context.Get("trace")
	if exists {
		if tc, ok := traceContext.(*util.TContext); ok {
			return tc
		}
	}
	return util.NewTrace()
}
