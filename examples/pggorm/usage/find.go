package main

import (
	"context"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/galaxy"
	"github.com/igordth/database-simplify/pggorm/usage/with"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func findGalaxy(ctx context.Context, gate *galaxy.Galaxy, log *zap.Logger) error {
	// SELECT * FROM "galaxies" ORDER BY "galaxies"."id" LIMIT 1
	gModel, err := gate.Find.One.Execute(ctx)
	if err != nil {
		return errors.Wrap(err, "Execute.One.Execute")
	}
	log.Debug("galaxy find one", zap.Reflect("model", gModel))

	// SELECT * FROM "galaxies" WHERE "galaxies"."id" = 2 ORDER BY "galaxies"."id" LIMIT 1
	gModel, err = gate.Find.One.Execute(ctx, 2)
	if err != nil {
		return errors.Wrap(err, "Execute.One.Execute with PK")
	}
	log.Debug("galaxy find one with PK=2", zap.Reflect("model", gModel))

	// SELECT * FROM "galaxies" WHERE name like ('Milky%') ORDER BY name desc LIMIT 5
	gModels, err := gate.Find.Many.
		With(
			with.Where("name like ?", "Milky%"),
			with.Order("name desc"),
		).
		With(with.Limit(5, 0)).
		Execute(ctx)
	if err != nil {
		return errors.Wrap(err, "Execute.Many.Execute")
	}
	log.Debug("galaxy find all", zap.Reflect("models", gModels))

	return nil
}
