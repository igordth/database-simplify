package pggorm

import "errors"

var (
	ErrSetupHost = errors.New("setup host in database config")
	ErrSetupName = errors.New("setup name in database config")
)
