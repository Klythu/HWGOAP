package database

import (
	"errors"
)

var (
	Err_conflict  = errors.New("Some things are not right")
	Err_not_found = errors.New("not found")
)
