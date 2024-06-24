package app

import (
	"github.com/adamnasrudin03/go-skeleton-gin/app/configs"
	"github.com/adamnasrudin03/go-skeleton-gin/app/controller"
	"github.com/adamnasrudin03/go-skeleton-gin/app/repository"
	"github.com/adamnasrudin03/go-skeleton-gin/app/service"
	"github.com/adamnasrudin03/go-skeleton-gin/pkg/driver"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB, cache *driver.RedisClient, cfg *configs.Configs, logger *logrus.Logger) *repository.Repositories {
	return &repository.Repositories{}
}

func WiringService(repo *repository.Repositories, cfg *configs.Configs, logger *logrus.Logger) *service.Services {
	return &service.Services{}
}

func WiringController(srv *service.Services, cfg *configs.Configs, logger *logrus.Logger) *controller.Controllers {
	return &controller.Controllers{}
}
