package dao

import "github.com/pkg/errors"

var (
	FailInsertError = errors.New("insert into table error")
	SelectError     = errors.New("select table error")
	NoneInsertError = errors.New("none inserted error")
)
