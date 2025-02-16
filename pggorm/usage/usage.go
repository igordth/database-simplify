package usage

import (
	"context"
	"github.com/igordth/database-simplify/pggorm"
	"github.com/igordth/database-simplify/pggorm/usage/with"
	"gorm.io/gorm"
)

type usage struct {
	pggorm.Connect
	ww []with.With
}

func (u *usage) prepareTx(ctx context.Context, model any) (tx *gorm.DB) {
	tx = u.Gorm(ctx).WithContext(ctx).Model(model)
	with.Apply(tx, u.ww...)
	return
}
