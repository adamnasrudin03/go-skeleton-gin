package repository

import (
	"github.com/adamnasrudin03/go-skeleton-gin/app/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TeamMemberRepository interface {
}

type TeamMemberRepo struct {
	DB     *gorm.DB
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewTeamMemberRepository(
	db *gorm.DB,
	cfg *configs.Configs,
	logger *logrus.Logger,
) TeamMemberRepository {
	return &TeamMemberRepo{
		DB:     db,
		Cfg:    cfg,
		Logger: logger,
	}
}
