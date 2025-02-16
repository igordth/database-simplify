package usage

import (
	"context"
	"github.com/igordth/database-simplify/pggorm"
	"github.com/igordth/database-simplify/pggorm/usage/with"
)

func NewUpdateCompare[T any](cnn pggorm.Connect) UpdateCompare[T] {
	return UpdateCompare[T]{
		Update: struct {
			Column  Update[*T]
			Columns Updates[*T]
		}{
			Column:  NewUpdate[*T](cnn),
			Columns: NewUpdates[*T](cnn),
		},
	}
}

type UpdateCompare[T any] struct {
	Update struct {
		Column  Update[*T]
		Columns Updates[*T]
	}
}

func NewUpdate[T any](cnn pggorm.Connect) Update[T] {
	return Update[T]{usage{Connect: cnn}}
}

type Update[T any] struct{ usage }

// Execute - updates column with value using callbacks
// [docs]: https://gorm.io/docs/update.html#Update-single-column
func (c *Update[T]) Execute(ctx context.Context, column string, value any) (rowsAffected int64, err error) {
	tx := c.prepareTx(ctx, new(T)).Update(column, value)
	return tx.RowsAffected, tx.Error
}

func (c *Update[T]) With(ww ...with.With) *Update[T] {
	c.ww = append(c.ww, ww...)
	return c
}

func NewUpdates[T any](cnn pggorm.Connect) Updates[T] {
	return Updates[T]{usage{Connect: cnn}}
}

type Updates[T any] struct{ usage }

// Execute - updates attributes using callbacks. values must be a struct or map
// [docs]: https://gorm.io/docs/update.html#Updates-multiple-columns
func (c *Updates[T]) Execute(ctx context.Context, values any) (rowsAffected int64, err error) {
	tx := c.prepareTx(ctx, new(T)).Updates(values)
	return tx.RowsAffected, tx.Error
}

func (c *Updates[T]) With(ww ...with.With) *Updates[T] {
	c.ww = append(c.ww, ww...)
	return c
}
