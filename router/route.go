package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raven0520/wallet/app"
	"github.com/raven0520/wallet/middleware"
)

// Init Router Initialization
func InitRouter() *gin.Engine {
	gin.SetMode(app.GetDebugMode())
	Router := gin.New()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	Router.Use(middleware.Cors())               // cross-domain
	Router.Use(middleware.TLSHandler())         // Handle HTTPS
	Router.NoRoute(middleware.ResponseNotFound) // return 404 if route not support
	Group := Router.Group("")                   // Default Route Group
	// Handle OPTIONS
	Group.OPTIONS("/:path/:method", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Options Request!")
	})
	// 基础路由
	Router.StaticFile("./favicon.ico", "./favicon.ico")
	Group.GET("/", func(context *gin.Context) {
		middleware.ResponseSuccess(context, "", "Request Success !")
	})
	Group.Use(
		middleware.RecoveryMiddleware(),
		middleware.TranslationMiddleware(),
	)
	BTCRegister(Group) // Register BTC route
	return Router
}
