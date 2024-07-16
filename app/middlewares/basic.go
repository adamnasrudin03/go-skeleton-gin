package middlewares

import (
	"net/http"

	"github.com/adamnasrudin03/go-skeleton-gin/app/configs"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func SetAuthBasic() gin.HandlerFunc {
	cfg := configs.GetInstance()
	return func(ctx *gin.Context) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := ctx.Request.BasicAuth()
		isValid := hasAuth && user == cfg.App.BasicUsername && password == cfg.App.BasicPassword
		if !isValid {
			err := helpers.NewError(helpers.ErrUnauthorized, helpers.NewResponseMultiLang(
				helpers.MultiLanguages{
					ID: "Token tidak valid",
					EN: "Invalid token",
				},
			))

			helpers.RenderJSON(ctx.Writer, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
