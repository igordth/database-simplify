package main

import (
	"context"
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/planet"
	"github.com/igordth/database-simplify/pggorm/usage/with"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func createPlanets(ctx context.Context, gate *planet.Planet, log *zap.Logger) error {
	// INSERT INTO "planets" ("name") VALUES ('Jupiter') RETURNING "id"
	model, _, err := gate.Create.One.
		With(with.Select("name")).
		Execute(ctx, &planet.Model{Name: "Jupiter"})
	if err != nil {
		return errors.Wrap(err, "Create.One.Execute")
	}
	log.Debug("planet create", zap.Reflect("model", model))

	// INSERT INTO "planets" ("name") VALUES ('Saturn'),('Uranus') RETURNING "id"
	models, RowsAffected, err := gate.Create.Many.
		With(with.Omit("star_id")).
		Execute(ctx, []planet.Model{
			{Name: "Saturn"},
			{Name: "Uranus"},
		})
	if err != nil {
		return errors.Wrap(err, "Create.Many.Execute")
	}
	log.Debug("planets create", zap.Reflect("models", models), zap.Int64("RowsAffected", RowsAffected))

	// INSERT INTO "planets" ("name") VALUES (initcap('neptune')) RETURNING "id"
	RowsAffected, err = gate.Create.Map.Execute(ctx, map[string]any{
		"name": gorm.Expr("initcap('neptune')"),
	})
	if err != nil {
		return errors.Wrap(err, "Create.Map.Execute")
	}
	log.Debug("planet create by map", zap.Int64("RowsAffected", RowsAffected))

	return nil
}
