package router

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/raven0520/wallet/app"
	"github.com/raven0520/wallet/middleware"
)

// BTCRegister Register BTC Router
func BTCRegister(group *gin.RouterGroup) {
	// BTC
	group.POST("/btc/:method", func(context *gin.Context) {
		c, connect, err := app.GetBTC()
		if err != nil {
			middleware.ResponseError(context, "unknown", err)
			return
		}
		defer connect.Close()
		Call(reflect.ValueOf(c), context)
	})
}
