package router

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type routes struct {
	HttpServer *gin.Engine
}

func NewRoutes() routes {
	var err error
	r := routes{
		HttpServer: gin.Default(),
	}

	r.HttpServer.Use(gin.Logger())
	r.HttpServer.Use(gin.Recovery())
	r.HttpServer.Use(cors.Default())

	r.HttpServer.GET("/", func(c *gin.Context) {
		response_mapper.RenderJSON(c.Writer, http.StatusOK, response_mapper.MultiLanguages{
			ID: "selamat datang di server ini",
			EN: "welcome this server",
		})
	})

	r.HttpServer.NoRoute(func(c *gin.Context) {
		err = response_mapper.ErrRouteNotFound()
		response_mapper.RenderJSON(c.Writer, http.StatusNotFound, err)
	})
	return r
}

func (r routes) Run(addr string) error {
	return r.HttpServer.Run(addr)
}
