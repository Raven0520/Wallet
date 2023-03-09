package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/raven0520/wallet/app"
	"github.com/unrolled/secure"
)

// TlsHandler Handle HTTPS
func TLSHandler() gin.HandlerFunc {
	var (
		host = app.GetStringConfig("base.http.host")
		port = app.GetStringConfig("base.http.port")
		addr = host + ":" + port
	)
	return func(context *gin.Context) {
		s := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     addr,
		})
		err := s.Process(context.Writer, context.Request)
		if err != nil {
			return
		}
		context.Next()
	}
}
