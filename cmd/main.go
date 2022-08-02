package main

import (
	"context"
	"fmt"
	"time"

	"net/http"
	"os"
	"os/signal"

	// "bitbucket.org/klopos/svc-inventory/conf"
	// "bitbucket.org/klopos/svc-inventory/pkg/db"
	// "bitbucket.org/klopos/svc-inventory/pkg/fflag"
	// "bitbucket.org/klopos/svc-inventory/pkg/mkafka"
	// "bitbucket.org/klopos/svc-inventory/pkg/mlog"
	// "bitbucket.org/klopos/svc-inventory/pkg/mvalidator"
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/beta-team/conf"
	"github.com/upgradeskill/beta-team/pkg/db"
	"github.com/upgradeskill/beta-team/pkg/mlog"
	"github.com/upgradeskill/beta-team/pkg/mvalidator"
	_ "go.uber.org/automaxprocs"
	"gorm.io/gorm"
)

const version = "1.0.0"
const appName = "minipos"

type app struct {
	echo      *echo.Echo
	cfg       *conf.Group
	logger    mlog.Logger
	database  *gorm.DB
	validator mvalidator.Validator
}

func main() {
	// load config
	cfg := conf.LoadConfig()

	// load logger
	logger := mlog.New("info", "stdout")

	// init database
	database := db.NewDatabase(&cfg.Database)

	// init validator
	mValidator := mvalidator.New()

	// create echo app
	e := echo.New()
	e.HideBanner = true

	// fill app
	application := app{
		echo:      e,
		cfg:       cfg,
		logger:    logger,
		database:  database,
		validator: mValidator,
	}

	// set common middleware
	// fulfill dependency and routing
	//PrepareCommonMiddleware(&application)
	PosRoute(&application)

	// Start echo server on goroutine
	go func() {
		e.Server.ReadTimeout = 45 * time.Second
		e.Server.WriteTimeout = 45 * time.Second
		e.Server.IdleTimeout = time.Minute

		if err := e.Start(fmt.Sprintf("0.0.0.0:%v", cfg.Pos.HTTPPort)); err != nil && err != http.ErrServerClosed {
			logger.Info("shutting down the server")
			panic(fmt.Sprintf("echo server startup panic: %s", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// gracefull shutdown stage ===============================================

	logger.Info("shutdown echo server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	// cleanup app ...
	logger.Info("Running cleanup tasks...")

	// close database
	db, _ := database.DB()
	db.Close()

	logger.Info("Done cleanup tasks...")
	logger.Sync()
}
