package planet

import (
	"github.com/igordth/database-simplify/pggorm"
	"github.com/igordth/database-simplify/pggorm/usage"
)

type Planet struct {
	pggorm.Connect
	usage.FindCompare[Model]
	usage.CountCompare[Model]
	usage.CreateCompare[Model]
	usage.SaveCompare[Model]
	usage.DeleteCompare[Model]
	usage.UpdateCompare[Model]
}

func New(cnn pggorm.Connect) *Planet {
	return &Planet{
		cnn,
		usage.NewFindCompare[Model](cnn),
		usage.NewCountCompare[Model](cnn),
		usage.NewCreateCompare[Model](cnn),
		usage.NewSaveCompare[Model](cnn),
		usage.NewDeleteCompare[Model](cnn),
		usage.NewUpdateCompare[Model](cnn),
	}
}
