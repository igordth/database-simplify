package usage

import (
	"context"
	"github.com/igordth/database-simplify/pggorm"
	"github.com/igordth/database-simplify/pggorm/usage/with"
)

func NewCountCompare[T any](cnn pggorm.Connect) CountCompare[T] {
	return CountCompare[T]{NewCount[T](cnn)}
}

type CountCompare[T any] struct {
	Count Count[T]
}

func NewCount[T any](cnn pggorm.Connect) Count[T] {
	return Count[T]{usage{Connect: cnn}}
}

type Count[T any] struct{ usage }

// Execute - get count of records
// [docs]: https://gorm.io/docs/advanced_query.html#Count
func (c *Count[T]) Execute(ctx context.Context) (cnt int64, err error) {
	err = c.prepareTx(ctx, new(T)).Count(&cnt).Error
	return
}

func (c *Count[T]) With(ww ...with.With) *Count[T] {
	c.ww = append(c.ww, ww...)
	return c
}
