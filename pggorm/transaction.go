package pggorm

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

const ctxTrx = "gorm-trx"

type Transaction interface {
	TrxBegin(ctx context.Context, opts ...*sql.TxOptions) context.Context
	TrxRollback(ctx context.Context)
	TrxCommit(ctx context.Context)
}

func (gw *Connection) TrxBegin(ctx context.Context, opts ...*sql.TxOptions) context.Context {
	return context.WithValue(ctx, ctxTrx, gw.gorm.Begin(opts...))
}

func (gw *Connection) TrxRollback(ctx context.Context) {
	trx := gw.getTrx(ctx)
	if trx == nil {
		return
	}
	trx.Rollback()
}

func (gw *Connection) TrxCommit(ctx context.Context) {
	trx := gw.getTrx(ctx)
	if trx == nil {
		return
	}
	trx.Commit()
}

func (gw *Connection) getTrx(ctx context.Context) *gorm.DB {
	tmp, _ := ctx.Value(ctxTrx).(*gorm.DB)
	return tmp
}
