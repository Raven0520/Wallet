package middleware

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raven0520/wallet/app"
	"github.com/raven0520/wallet/util"
)

// RecoveryMiddleware catch all panic, then return error json
func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if app.BaseConf.Base.DebugMode != "debug" {
					ResponseSystemError(context, util.ErrSystem)
					return
				}
				ResponseSystemError(context, errors.New(fmt.Sprint(err)))
				return
			}
		}()
		context.Next()
	}
}
