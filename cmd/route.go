package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/beta-team/pkg/web"
)

func PosRoute(app *app) {
	// init dependency
	//userRepo := web.NewUserRepo(app.database)

	// iapRepository := cogsrepo.NewIAPRepository(app.database)
	// historyRepository := cogsrepo.NewHistoryRepository(app.database)
	// summaryRepository := cogsrepo.NewSummaryRepository(app.database)
	// itemBranchRepository := itembranch.NewItemRepository(app.kloposDatabase) // use klopos database
	// kafkaPublisherRepo := cogsrepo.NewKafkaPublisher(app.kafkaProducer, &app.cfg.Kafka, app.logger)

	// eventjoin.RunInBackground(kafkaPublisherRepo, app.logger)
	// eventJoiner := eventjoin.NewEventJoiner()

	//marketplaceAggrClient := cogsrepo.NewMarketplaceAggrHttpClient(app.cfg, app.logger)

	// cogsUsecase := cogscore.NewCogsCore(
	// 	app.logger,
	// 	app.cfg,
	// 	app.ffClient,
	// 	iapRepository,
	// 	historyRepository,
	// 	summaryRepository,
	// 	kafkaPublisherRepo,
	// 	eventJoiner,
	// 	marketplaceAggrClient,
	// 	itemBranchRepository,
	// )

	// create handler
	//cogsHandler := handler.NewCogsHandler(cogsUsecase, app.validator, app.logger, app.ffClient, *app.cfg)

	// init additional middleware here or directly in route (ex. JWT, api key)
	// ...

	// set route =============================================================
	// route for check health
	app.echo.GET("v1/health/ping", func(c echo.Context) error {
		return web.ResponseFormatter(c, http.StatusOK, "Success", map[string]any{"status": "ok"}, nil)
	})

	//g := app.echo.Group("/v1/cogs")

	// with jwt key
	//g.GET("/recalculate-status/:branch_id", cogsHandler.RecalculateStatusByBranchID, mid.JwtMiddleware(app.cfg.Secrets.MayangSecretKey))

	// with API key Middleware
	// g.Use(mid.ApiKeyMiddleware(app.cfg.Secrets.CogsApiKey))
	// g.GET("", cogsHandler.Get)
	// g.PUT("/update", cogsHandler.Update)
	// g.PUT("/edit", cogsHandler.Edit)
	// g.POST("/void", cogsHandler.Delete)
	// g.POST("/recalculate", cogsHandler.RecalculateByDate)
	// g.POST("/recalculate/branch", cogsHandler.RecalculateBranch)
}
