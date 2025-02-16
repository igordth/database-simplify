package usage

import (
	"context"
	"github.com/igordth/database-simplify/pggorm"
	"github.com/igordth/database-simplify/pggorm/usage/with"
)

func NewCreateCompare[T any](cnn pggorm.Connect) CreateCompare[T] {
	return CreateCompare[T]{
		Create: struct {
			One  CreateModel[*T]
			Many CreateModel[[]T]
			Map  CreateMap[*T]
		}{
			One:  NewCreateModel[*T](cnn),
			Many: NewCreateModel[[]T](cnn),
			Map:  NewCreateMap[*T](cnn),
		},
	}
}

type CreateCompare[T any] struct {
	Create struct {
		One  CreateModel[*T]
		Many CreateModel[[]T]
		Map  CreateMap[*T]
	}
}

func NewCreateModel[T any](cnn pggorm.Connect) CreateModel[T] {
	return CreateModel[T]{usage{Connect: cnn}}
}

type CreateModel[T any] struct{ usage }

// Execute - create record by model
// [docs]: https://gorm.io/docs/create.html#Create-Record
func (c *CreateModel[T]) Execute(ctx context.Context, model T) (m T, rowsAffected int64, err error) {
	tx := c.prepareTx(ctx, model).Create(model)
	return model, tx.RowsAffected, tx.Error
}

func (c *CreateModel[T]) With(ww ...with.With) *CreateModel[T] {
	c.ww = append(c.ww, ww...)
	return c
}

func NewCreateMap[T any](cnn pggorm.Connect) CreateMap[T] {
	return CreateMap[T]{usage{Connect: cnn}}
}

type CreateMap[T any] struct{ usage }

// Execute - create record by map
// [docs]: https://gorm.io/docs/create.html#Create-Record
func (c *CreateMap[T]) Execute(ctx context.Context, data map[string]any) (rowsCreated int64, err error) {
	tx := c.prepareTx(ctx, new(T)).Create(data)
	return tx.RowsAffected, tx.Error
}

func (c *CreateMap[T]) With(ww ...with.With) *CreateMap[T] {
	c.ww = append(c.ww, ww...)
	return c
}
