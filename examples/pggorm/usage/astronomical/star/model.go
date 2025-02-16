package star

import (
	"github.com/igordth/database-simplify/examples/pggorm/usage/astronomical/galaxy"
)

type Model struct {
	ID       uint   `gorm:"primary_key;column:id"`
	Name     string `gorm:"column:name"`
	GalaxyId *int   `gorm:"column:galaxy_id"`
	Galaxy   *galaxy.Model
}

func (Model) TableName() string {
	return "stars"
}
