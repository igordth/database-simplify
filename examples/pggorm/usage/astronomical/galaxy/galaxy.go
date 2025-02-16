package galaxy

import (
	"github.com/igordth/database-simplify/pggorm"
	"github.com/igordth/database-simplify/pggorm/usage"
)

type Galaxy struct {
	pggorm.Connect
	usage.FindCompare[Model]
}

func New(cnn pggorm.Connect) *Galaxy {
	return &Galaxy{
		cnn,
		usage.NewFindCompare[Model](cnn),
	}
}
