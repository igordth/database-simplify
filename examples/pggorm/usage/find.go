package main

import (
	"context"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/galaxy"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/planet"
	"github.com/igordth/database-simplify/pggorm/usage/with"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func findGalaxy(ctx context.Context, gate *galaxy.Galaxy, log *zap.Logger) error {
	// SELECT * FROM "galaxies" ORDER BY "galaxies"."id" LIMIT 1
	model, err := gate.Find.One.Execute(ctx)
	if err != nil {
		return errors.Wrap(err, "Find.One.Execute")
	}
	log.Debug("galaxy find one", zap.Reflect("model", model))

	// SELECT * FROM "galaxies" WHERE "galaxies"."id" = 2 ORDER BY "galaxies"."id" LIMIT 1
	model, err = gate.Find.One.Execute(ctx, 2)
	if err != nil {
		return errors.Wrap(err, "Find.One.Execute with PK")
	}
	log.Debug("galaxy find one with PK=2", zap.Reflect("model", model))

	// SELECT * FROM "galaxies" WHERE name like 'Milky%' ORDER BY name desc LIMIT 5
	models, err := gate.Find.Many.
		With(
			with.Where("name like ?", "Milky%"),
			with.Order("name desc"),
		).
		With(with.Limit(5, 0)).
		Execute(ctx)
	if err != nil {
		return errors.Wrap(err, "Find.Many.Execute")
	}
	log.Debug("galaxy find all", zap.Reflect("models", models))

	return nil
}

func findPlanetsWithPreload(ctx context.Context, gate *planet.Planet, log *zap.Logger) error {
	// SELECT * FROM "stars" WHERE "stars"."id" = 1
	// SELECT * FROM "planets" ORDER BY "planets"."id" LIMIT 1
	model, err := gate.Find.One.With(with.Preload("Star")).Execute(ctx)
	if err != nil {
		return errors.Wrap(err, "Find.One.Execute")
	}
	log.Debug("planet find one with preload", zap.Reflect("model", model))

	// SELECT * FROM "stars" WHERE "stars"."id" = 1
	// SELECT * FROM "planets"
	models, err := gate.Find.Many.With(with.Preload("Star")).Execute(ctx)
	if err != nil {
		return errors.Wrap(err, "Find.Many.Execute")
	}
	log.Debug("planet find all with preload", zap.Reflect("models", models))

	return nil
}
