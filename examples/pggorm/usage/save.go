package main

import (
	"context"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/planet"
	"github.com/igordth/database-simplify/pggorm/usage/with"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func savePlanet(ctx context.Context, gate *planet.Planet, log *zap.Logger) error {
	// UPDATE "planets" SET "star_id"=1 WHERE "id" = 5
	sunId := 1
	model, err := gate.Save.
		With(with.Select("star_id")).
		Execute(ctx, &planet.Model{ID: 5, StarId: &sunId})
	if err != nil {
		return errors.Wrap(err, "Save.Execute")
	}
	log.Debug("save planet", zap.Reflect("model", model))

	return nil
}
