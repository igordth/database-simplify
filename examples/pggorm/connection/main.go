package main

import (
	"context"
	"github.com/igordth/database-simplify/pggorm"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	// zap logger
	zapLog, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// connection
	cnn, df, err := pggorm.NewConnection(
		gormConfig,
		pggorm.NewLog(zapLog, logConfig),
	)
	if err != nil {
		zapLog.Panic("pggorm.NewConnection", zap.Error(err))
	}
	defer df()

	// SOME QUERIES EXAMPLE

	// find all from table galaxies
	var galaxies []Galaxy
	err = cnn.Gorm(ctx).Find(&galaxies).Error
	if err != nil {
		zapLog.Panic("find galaxies", zap.Error(err))
	}
	zapLog.Debug("find galaxies", zap.Reflect("galaxies", galaxies))

	// find all stars and preload galaxy
	var stars []Stars
	err = cnn.Gorm(ctx).Preload("Galaxy").Find(&stars).Error
	if err != nil {
		zapLog.Panic("find stars", zap.Error(err))
	}
	zapLog.Debug("find stars with galaxy", zap.Reflect("stars", stars))

	// find all from table planets
	var planets []Planet
	err = cnn.Gorm(ctx).Preload("Star").Find(&planets).Error
	if err != nil {
		zapLog.Panic("find planets", zap.Error(err))
	}
	zapLog.Debug("find planets with star", zap.Reflect("planets", planets))
}
