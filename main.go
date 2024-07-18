package main

import (
	"fmt"

	"github.com/adamnasrudin03/go-skeleton-gin/app"
	"github.com/adamnasrudin03/go-skeleton-gin/app/router"
	"github.com/adamnasrudin03/go-skeleton-gin/configs"
	"github.com/adamnasrudin03/go-skeleton-gin/pkg/database"
	"github.com/adamnasrudin03/go-skeleton-gin/pkg/driver"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func main() {
	configs.LoadEnv()
	var (
		cfg                  = configs.GetInstance()
		logger               = driver.Logger(cfg)
		cache                = driver.Redis(cfg)
		validate             = validator.New()
		db          *gorm.DB = database.SetupDbConnection(cfg, logger)
		repo                 = app.WiringRepository(db, &cache, cfg, logger)
		services             = app.WiringService(repo, cfg, logger)
		controllers          = app.WiringController(services, cfg, logger, validate)
	)

	defer database.CloseDbConnection(db, logger)

	r := router.NewRoutes(*controllers)

	listen := fmt.Sprintf(":%v", cfg.App.Port)
	r.Run(listen)
}
