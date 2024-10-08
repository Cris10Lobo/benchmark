package webframeworks

import (
	"log"
	"net/http"
	"testing"

	"github.com/lkumarjain/benchmark/web-frameworks/gin"
	"github.com/lkumarjain/benchmark/web-frameworks/routes"
)

var routeTests = []routes.Routes{
	{
		Method: http.MethodGet,
		Path:   "/ping",
	},
}

func Test_Frameworks(suite *testing.T) {

	suite.Run("Gin", func(t *testing.T) {
		svr := gin.Handler(routeTests)
		http.ListenAndServe(":8080", svr)

		log.Println("Server started at :8080")
	})

}
