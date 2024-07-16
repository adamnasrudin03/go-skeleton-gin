package router

import (
	"github.com/adamnasrudin03/go-skeleton-gin/app/controller"
	"github.com/adamnasrudin03/go-skeleton-gin/app/middlewares"
	"github.com/gin-gonic/gin"
)

func (r routes) tmRouter(rg *gin.RouterGroup, handler controller.TeamMemberController) {
	tm := rg.Group("/team-members")
	{
		tm.POST("/", middlewares.SetAuthBasic(), handler.Create)
		tm.GET("/:id", handler.GetDetail)
		tm.DELETE("/:id", middlewares.SetAuthBasic(), handler.Delete)
	}

}
