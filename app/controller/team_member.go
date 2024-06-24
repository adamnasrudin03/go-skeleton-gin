package controller

import (
	"github.com/adamnasrudin03/go-skeleton-gin/app/service"
	"github.com/sirupsen/logrus"
)

type TeamMemberController interface {
}

type TMemberController struct {
	Service service.TeamMemberService
	Logger  *logrus.Logger
}

func NewTeamMemberDelivery(
	srv service.TeamMemberService,
	logger *logrus.Logger,
) TeamMemberController {
	return &TMemberController{
		Service: srv,
		Logger:  logger,
	}
}
