package db

import (
	"fmt"
	"log"
	"time"

	"github.com/upgradeskill/beta-team/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(config *conf.Database) *gorm.DB {
	// gorm mysql orm
	dbConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Schema,
	)

	maxConn := 100
	if config.MaxConn != 0 {
		maxConn = config.MaxConn
	}
	return createDatabaseInstance(dbConfig, maxConn)
}

func createDatabaseInstance(dsn string, maxConn int) *gorm.DB {
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		// Logger:                 logger.Default.LogMode(logger.Silent),
	}

	inst, err := gorm.Open(
		mysql.Open(dsn),
		gormConfig,
	)
	if err != nil {
		log.Panic(err)
	}

	db, err := inst.DB()
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(maxConn)
	db.SetConnMaxLifetime(time.Minute * 20)

	return inst
}
