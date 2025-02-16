package usage

import (
	"context"
	"github.com/igordth/database-simplify/pggorm"
	"github.com/igordth/database-simplify/pggorm/usage/with"
)

func NewDeleteCompare[T any](cnn pggorm.Connect) DeleteCompare[T] {
	return DeleteCompare[T]{
		Delete: NewDelete[*T](cnn),
	}
}

type DeleteCompare[T any] struct {
	Delete Delete[*T]
}

func NewDelete[T any](cnn pggorm.Connect) Delete[T] {
	return Delete[T]{usage{Connect: cnn}}
}

type Delete[T any] struct{ usage }

// Execute - deletes value matching given conditions
// [docs]: https://gorm.io/docs/delete.html
func (c *Delete[T]) Execute(ctx context.Context, value any, conds ...any) (err error) {
	return c.prepareTx(ctx, new(T)).Delete(value, conds).Error
}

func (c *Delete[T]) With(ww ...with.With) *Delete[T] {
	c.ww = append(c.ww, ww...)
	return c
}
