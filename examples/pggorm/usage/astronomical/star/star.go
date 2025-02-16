package star

import (
	"github.com/igordth/database-simplify/pggorm"
)

type Star struct {
	pggorm.Connect
}

func New(cnn pggorm.Connect) *Star {
	return &Star{cnn}
}
