package dao

import "github.com/pkg/errors"

var (
	InsertError        = errors.New("insert into table error")
	SelectError        = errors.New("select table error")
	NothingInsertError = errors.New("nothing inserted error")
	NothingSelectError = errors.New("nothing select error")
)
