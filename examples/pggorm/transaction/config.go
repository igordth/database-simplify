package main

import (
	"github.com/igordth/database-simplify/pggorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	gormConfig = pggorm.Config{
		Name:        "astronomical_objects",
		User:        "user",
		Password:    "password",
		Host:        "localhost",
		MaxOpenConn: 1,
		MaxIdleConn: 1,
	}
	logConfig = logger.Config{
		SlowThreshold:             time.Second,
		Colorful:                  true,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Info,
	}
)
