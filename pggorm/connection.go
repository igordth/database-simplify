package pggorm

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connect interface {
	DB() *sql.DB
	Gorm(ctx context.Context) *gorm.DB
	Transaction
}

type Connection struct {
	gorm *gorm.DB
	db   *sql.DB
}

func NewLog(log *zap.Logger, cfg logger.Config) logger.Interface {
	if log == nil {
		return logger.Discard
	}
	return logger.New(zap.NewStdLog(log), cfg)
}

func NewConnection(cfg Config, log logger.Interface) (c Connect, df func(), err error) {
	ds, err := cfg.String()
	if err != nil {
		return
	}
	grm, err := gorm.Open(postgres.Open(ds), &gorm.Config{Logger: log})
	if err != nil {
		err = errors.Wrap(err, "gorm.Open")
		return
	}
	db, err := grm.DB()
	if err != nil {
		err = errors.Wrap(err, "try get DB")
		return
	}
	df = func() { _ = db.Close() }
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	return &Connection{gorm: grm, db: db}, df, nil
}

func (gw *Connection) DB() *sql.DB {
	return gw.db
}

func (gw *Connection) Gorm(ctx context.Context) *gorm.DB {
	if trx := gw.getTrx(ctx); trx != nil {
		return trx
	}
	return gw.gorm
}
