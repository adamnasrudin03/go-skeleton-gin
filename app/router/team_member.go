package router

import (
	"net/http"

	"github.com/adamnasrudin03/go-skeleton-gin/app/controller"
	gtHelpers "github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (r routes) tmRouter(rg *gin.RouterGroup, handler controller.TeamMemberController) {
	tm := rg.Group("/team-members")
	{
		tm.GET("/", func(c *gin.Context) {
			gtHelpers.RenderJSON(c.Writer, http.StatusOK, "Build with love by adamnasrudin03")
		})
	}

}
