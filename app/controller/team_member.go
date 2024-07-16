package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/adamnasrudin03/go-skeleton-gin/app/dto"
	"github.com/adamnasrudin03/go-skeleton-gin/app/service"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type TeamMemberController interface {
	Create(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetList(ctx *gin.Context)
}

type TMemberController struct {
	Service  service.TeamMemberService
	Logger   *logrus.Logger
	Validate *validator.Validate
}

func NewTeamMemberDelivery(
	srv service.TeamMemberService,
	logger *logrus.Logger,
	validator *validator.Validate,
) TeamMemberController {
	return &TMemberController{
		Service:  srv,
		Logger:   logger,
		Validate: validator,
	}
}

func (c *TMemberController) Create(ctx *gin.Context) {
	var (
		opName = "TeamMemberController-Create"
		input  dto.TeamMemberCreateReq
		err    error
	)

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	// validation input user
	err = c.Validate.Struct(input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.FormatValidationError(err))
		return
	}

	res, err := c.Service.Create(ctx, input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusCreated, res)
}

func (c *TMemberController) GetDetail(ctx *gin.Context) {
	var (
		opName  = "TeamMemberController-GetDetail"
		idParam = strings.TrimSpace(ctx.Param("id"))
		err     error
	)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.Logger.Errorf("%v error parse param: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrInvalid("ID Anggota team", "Team Member ID"))
		return
	}

	res, err := c.Service.GetByID(ctx, id)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, res)
}

func (c *TMemberController) Delete(ctx *gin.Context) {
	var (
		opName  = "TeamMemberController-Delete"
		idParam = strings.TrimSpace(ctx.Param("id"))
		err     error
	)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.Logger.Errorf("%v error parse param: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrInvalid("ID Anggota team", "Team Member ID"))
		return
	}

	err = c.Service.DeleteByID(ctx, id)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, "Team Member Deleted")
}

func (c *TMemberController) Update(ctx *gin.Context) {
	var (
		opName  = "TeamMemberController-Update"
		idParam = strings.TrimSpace(ctx.Param("id"))
		input   dto.TeamMemberUpdateReq
		err     error
	)

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.Logger.Errorf("%v error parse param: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrInvalid("ID Anggota team", "Team Member ID"))
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}
	input.ID = id
	// validation input user
	err = c.Validate.Struct(input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.FormatValidationError(err))
		return
	}

	err = c.Service.Update(ctx, input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, "Team Member Updated")
}

func (c *TMemberController) GetList(ctx *gin.Context) {
	var (
		opName = "TeamMemberController-GetList"
		input  dto.TeamMemberListReq
		err    error
	)

	err = ctx.ShouldBindQuery(&input)
	if err != nil {
		c.Logger.Errorf("%v error bind json: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.ErrGetRequest())
		return
	}

	// validation input user
	err = c.Validate.Struct(input)
	if err != nil {
		helpers.RenderJSON(ctx.Writer, http.StatusBadRequest, helpers.FormatValidationError(err))
		return
	}

	res, err := c.Service.GetList(ctx, input)
	if err != nil {
		c.Logger.Errorf("%v error: %v ", opName, err)
		helpers.RenderJSON(ctx.Writer, http.StatusInternalServerError, err)
		return
	}

	helpers.RenderJSON(ctx.Writer, http.StatusOK, res)
}
