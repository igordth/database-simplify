package main

import (
	"context"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/galaxy"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/planet"
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

	// Connection
	cnn, df, err := pggorm.NewConnection(
		gormConfig,
		pggorm.NewLog(zapLog, logConfig),
	)
	if err != nil {
		zapLog.Panic("pggorm.NewConnection", zap.Error(err))
	}
	defer df()

	// Define gates
	gl := galaxy.New(cnn)
	//st := star.New(cnn)
	pl := planet.New(cnn)

	// Work with find in galaxies
	if err = findGalaxy(ctx, gl, zapLog); err != nil {
		zapLog.Panic("findGalaxy", zap.Error(err))
	}

	// Work with count in planets
	if err = countPlanets(ctx, pl, zapLog); err != nil {
		zapLog.Panic("countPlanets", zap.Error(err))
	}

	// Work with create planets
	if err = createPlanets(ctx, pl, zapLog); err != nil {
		zapLog.Panic("countPlanets", zap.Error(err))
	}

	// Work with save planet
	if err = savePlanet(ctx, pl, zapLog); err != nil {
		zapLog.Panic("savePlanet", zap.Error(err))
	}

	// Work with update planet
	if err = updatePlanets(ctx, pl, zapLog); err != nil {
		zapLog.Panic("updatePlanets", zap.Error(err))
	}

	// Work with delete planet
	if err = deletePlanet(ctx, pl, zapLog); err != nil {
		zapLog.Panic("deletePlanet", zap.Error(err))
	}
}
