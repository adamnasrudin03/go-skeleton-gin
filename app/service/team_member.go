package service

import (
	"github.com/adamnasrudin03/go-skeleton-gin/app/configs"
	"github.com/adamnasrudin03/go-skeleton-gin/app/repository"
	"github.com/sirupsen/logrus"
)

type TeamMemberService interface {
}

type TeamMemberSrv struct {
	Repo   repository.TeamMemberRepository
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewTeamMemberService(
	tmRepo repository.TeamMemberRepository,
	cfg *configs.Configs,
	logger *logrus.Logger,
) TeamMemberService {
	return TeamMemberSrv{
		Repo:   tmRepo,
		Cfg:    cfg,
		Logger: logger,
	}
}
