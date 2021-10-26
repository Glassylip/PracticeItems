package model

import (
	"errors"
)

var (
	ERROR_USER_NOTEXISTS = errors.New("User does not exist..")
	ERROR_USER_EXISTS    = errors.New("User exit...")
	ERROR_USER_PWD       = errors.New("Wrong password")
)
