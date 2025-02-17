package usage

import (
	"context"
	"github.com/igordth/database-simplify/pggorm"
	"github.com/igordth/database-simplify/pggorm/usage/with"
	"reflect"
)

func NewFindCompare[T any](cnn pggorm.Connect) (c FindCompare[T]) {
	c = FindCompare[T]{}
	c.Find.One = NewFind[*T](cnn)
	c.Find.Many = NewFind[[]T](cnn)
	return
}

type FindCompare[T any] struct {
	Find struct {
		One  Find[*T]
		Many Find[[]T]
	}
}

func NewFind[T any](cnn pggorm.Connect) Find[T] {
	return Find[T]{usage{Connect: cnn}}
}

type Find[T any] struct{ usage }

// Execute - Retrieving object(s) with conditions
// [docs]: https://gorm.io/docs/query.html
func (fn *Find[T]) Execute(ctx context.Context, conds ...any) (T, error) {
	var res T
	tx := fn.prepareTx(ctx, res)
	switch reflect.TypeOf(*new(T)).Kind() {
	case reflect.Pointer:
		tx.First(&res, conds...)
	default:
		tx.Find(&res, conds...)
	}
	return res, tx.Error
}

func (fn *Find[T]) With(ww ...with.With) *Find[T] {
	fn.ww = append(fn.ww, ww...)
	return fn
}
