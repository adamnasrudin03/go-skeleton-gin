package controller

import (
	"net/http"
	"strconv"
	"strings"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-skeleton-gin/app/dto"
	"github.com/adamnasrudin03/go-skeleton-gin/app/middlewares"
	"github.com/adamnasrudin03/go-skeleton-gin/app/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type TeamMemberController interface {
	Mount(rg *gin.RouterGroup)
	Create(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetList(ctx *gin.Context)
}

type TeamMemberHandler struct {
	Service  service.TeamMemberService
	Logger   *logrus.Logger
	Validate *validator.Validate
}

func NewTeamMemberDelivery(
	srv service.TeamMemberService,
	logger *logrus.Logger,
	validator *validator.Validate,
) TeamMemberController {
	return &TeamMemberHandler{
		Service:  srv,
		Logger:   logger,
		Validate: validator,
	}
}

func (c *TeamMemberHandler) Mount(rg *gin.RouterGroup) {
	tm := rg.Group("/team-members")
	{
		tm.GET("", c.GetList)
		tm.POST("", middlewares.SetAuthBasic(), c.Create)
		tm.GET("/:id", c.GetDetail)
		tm.DELETE("/:id", middlewares.SetAuthBasic(), c.Delete)
		tm.PUT("/:id", middlewares.SetAuthBasic(), c.Update)
	}
}

func (c *TeamMemberHandler) getParamID(ctx *gin.Context) (uint64, error) {
	idParam := strings.TrimSpace(ctx.Param("id"))
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.Logger.Errorf("TeamMemberController-getParamID error parse param: %v ", err)
		return 0, response_mapper.ErrInvalid("ID Anggota team", "Team Member ID")
	}
	return id, nil
}

func (c *TeamMemberHandler) Create(ctx *gin.Context) {
	var (
		opName = "TeamMemberController-Create"
		input  dto.TeamMemberCreateReq
		err    error
	)

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	// validation input user
	err = c.Validate.Struct(input)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.FormatValidationError(err))
		return
	}

	res, err := c.Service.Create(ctx, input)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusCreated, res)
}

func (c *TeamMemberHandler) GetDetail(ctx *gin.Context) {
	var (
		opName = "TeamMemberController-GetDetail"
		err    error
	)

	id, err := c.getParamID(ctx)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, err)
		return
	}

	res, err := c.Service.GetByID(ctx, id)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, res)
}

func (c *TeamMemberHandler) Delete(ctx *gin.Context) {
	var (
		opName = "TeamMemberController-Delete"
		err    error
	)

	id, err := c.getParamID(ctx)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, err)
		return
	}

	err = c.Service.DeleteByID(ctx, id)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, response_mapper.MultiLanguages{
		ID: "Anggota Tim Berhasil Dihapus",
		EN: "Team Member Deleted Successfully",
	})
}

func (c *TeamMemberHandler) Update(ctx *gin.Context) {
	var (
		opName = "TeamMemberController-Update"
		input  dto.TeamMemberUpdateReq
		err    error
	)

	id, err := c.getParamID(ctx)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}
	input.ID = id
	// validation input user
	err = c.Validate.Struct(input)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.FormatValidationError(err))
		return
	}

	err = c.Service.Update(ctx, input)
	if err != nil {
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, response_mapper.MultiLanguages{
		ID: "Anggota Tim Berhasil Diperbarui",
		EN: "Team Member Updated Successfully",
	})
}

func (c *TeamMemberHandler) GetList(ctx *gin.Context) {
	var (
		opName = "TeamMemberController-GetList"
		input  dto.TeamMemberListReq
		err    error
	)

	err = ctx.ShouldBindQuery(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusBadRequest, response_mapper.ErrGetRequest())
		return
	}

	res, err := c.Service.GetList(ctx, input)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		response_mapper.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	response_mapper.RenderJSON(ctx.Writer, http.StatusOK, res)
}
