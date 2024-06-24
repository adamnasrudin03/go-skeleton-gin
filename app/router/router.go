package router

import (
	"net/http"

	"github.com/adamnasrudin03/go-skeleton-gin/app/controller"
	gtHelpers "github.com/adamnasrudin03/go-template/pkg/helpers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes(h controller.Controllers) routes {
	var err error
	r := routes{
		router: gin.Default(),
	}

	r.router.Use(gin.Logger())
	r.router.Use(gin.Recovery())
	r.router.Use(cors.Default())

	r.router.GET("/", func(c *gin.Context) {
		gtHelpers.RenderJSON(c.Writer, http.StatusOK, "welcome this server")
	})

	v1 := r.router.Group("/api/v1")
	r.tmRouter(v1, h.TeamMember)

	r.router.NoRoute(func(c *gin.Context) {
		err = gtHelpers.ErrRouteNotFound()
		gtHelpers.RenderJSON(c.Writer, http.StatusNotFound, err)
	})
	return r
}

func (r routes) Run(addr string) error {
	return r.router.Run(addr)
}
