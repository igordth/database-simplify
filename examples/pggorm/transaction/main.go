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

	// Work with transaction
	sunId := 1
	ctxTx := cnn.TrxBegin(ctx)
	// add new planet Pluto to planets of Sun in transaction
	err = cnn.Gorm(ctxTx).
		Create(&Planet{
			BaseModel: BaseModel{Name: "Pluto"},
			StarId:    &sunId,
		}).
		Error
	if err != nil {
		panic(err)
	}
	LogSunPlanets(ctxTx, cnn, zapLog) // Pluto - contain in planets of Sun
	cnn.TrxRollback(ctxTx)

	// Pluto - not contain in planets of Sun because using Rollback
	LogSunPlanets(ctx, cnn, zapLog)
}

func LogSunPlanets(ctx context.Context, cnn pggorm.Connect, log *zap.Logger) {
	var planets []Planet
	err := cnn.Gorm(ctx).Where("star_id", 1).Find(&planets).Error
	if err != nil {
		log.Panic("find planets with star_id=1[Sun]", zap.Error(err))
	}
	log.Debug("find planets with star_id=1[Sun]", zap.Reflect("planets", planets))
}
