package usage

import (
	"context"
	"github.com/igordth/database-simplify/pggorm"
	"github.com/igordth/database-simplify/pggorm/usage/with"
)

func NewSaveCompare[T any](cnn pggorm.Connect) SaveCompare[T] {
	return SaveCompare[T]{
		Save: NewSave[*T](cnn),
	}
}

type SaveCompare[T any] struct {
	Save Save[*T]
}

func NewSave[T any](cnn pggorm.Connect) Save[T] {
	return Save[T]{usage{Connect: cnn}}
}

type Save[T any] struct{ usage }

// Execute - updates value in database. If value doesn't contain a matching primary key, value is inserted.
// [docs]: https://gorm.io/docs/update.html#Save-All-Fields
func (c *Save[T]) Execute(ctx context.Context, model T) (m T, err error) {
	err = c.prepareTx(ctx, model).Save(model).Error
	return model, err
}

func (c *Save[T]) With(ww ...with.With) *Save[T] {
	c.ww = append(c.ww, ww...)
	return c
}
