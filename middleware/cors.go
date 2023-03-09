package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var allows = []string{"http://localhost", "http://localhost:3000", "https://console.raven.pub"}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		if origin == "" {
			origin = "http://localhost:3000"
		} // Default Origin
		for _, allow := range allows {
			if origin == allow {
				origin = allow
			}
		}
		context.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		context.Header("Access-Control-Allow-Origin", origin)
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// Header Type
		context.Header("Access-Control-Allow-Headers", "Authorization, UUID, Sign, Driver, Content-Length, X-CSRF-Token, token,Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "false")
		context.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

		// Pass All OPTIONS
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		context.Next()
	}
}
