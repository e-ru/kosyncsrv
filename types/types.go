package types

import "database/sql"

type DBApi interface {
	MustExec(query string, args ...interface{}) sql.Result
	Get(dest interface{}, query string, args ...interface{}) error
}
