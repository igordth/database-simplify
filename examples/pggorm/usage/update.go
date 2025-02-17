package main

import (
	"context"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/planet"
	"github.com/igordth/database-simplify/pggorm/usage/with"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func updatePlanets(ctx context.Context, gate *planet.Planet, log *zap.Logger) error {
	// UPDATE "planets" SET "star_id"=1 WHERE "id" = 6
	affected, err := gate.Update.Column.
		With(with.Where("id", 6)).
		Execute(ctx, "star_id", 1)
	if err != nil {
		return errors.Wrap(err, "Update.One.Execute")
	}
	log.Debug("update one planet", zap.Reflect("affected", affected))

	// UPDATE "planets" SET "star_id"=1 WHERE "id" IN (7,8,9)
	affected, err = gate.Update.Columns.
		With(with.Where(7, 8, 9)).
		Execute(ctx, map[string]any{
			"star_id": 1,
		})
	if err != nil {
		return errors.Wrap(err, "Update.One.Execute")
	}
	log.Debug("update one planet", zap.Reflect("affected", affected))

	return nil
}
