package planet

import (
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/star"
)

type Model struct {
	ID     uint   `gorm:"primary_key;column:id"`
	Name   string `gorm:"column:name"`
	StarId *int   `gorm:"column:star_id"`
	Star   *star.Model
}

func (Model) TableName() string {
	return "planets"
}
