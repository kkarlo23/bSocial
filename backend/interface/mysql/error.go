package mysql

import "errors"

func GenericDBError() error {
	return errors.New("database error")
}
