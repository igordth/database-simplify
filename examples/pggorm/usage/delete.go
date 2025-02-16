package main

import (
	"context"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/planet"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func deletePlanet(ctx context.Context, gate *planet.Planet, log *zap.Logger) error {
	//  DELETE FROM "planets" WHERE "planets"."id" = 9
	err := gate.Delete.Execute(ctx, "id", 9)
	if err != nil {
		return errors.Wrap(err, "Delete.Execute")
	}
	log.Debug("delete planet Neptune[9]")
	return nil
}
