package main

import (
	"context"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/planet"
	"github.com/igordth/database-simplify/pggorm/usage/with"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func countPlanets(ctx context.Context, gate *planet.Planet, log *zap.Logger) error {
	// SELECT count(*) FROM "planets"
	count, err := gate.Count.Execute(ctx)
	if err != nil {
		return errors.Wrap(err, "Execute.Execute")
	}
	log.Debug("count of all planets", zap.Reflect("count", count))

	// SELECT COUNT(DISTINCT("star_id")) FROM "planets"
	count, err = gate.
		Count.
		With(with.Distinct("star_id")).
		Execute(ctx)
	if err != nil {
		return errors.Wrap(err, "Execute.Execute")
	}
	log.Debug("count of all planets", zap.Reflect("count", count))

	return nil
}
