package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lkumarjain/benchmark/web-frameworks/routes"
)

func ginHandle(_ *gin.Context) {}

func Handler(routes []routes.Routes) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	svr := gin.New()

	for _, route := range routes {
		svr.Handle(route.Method, route.Path, ginHandle)
	}

	return svr
}
