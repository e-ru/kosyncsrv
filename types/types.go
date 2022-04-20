package types

import "database/sql"

// type DBApi interface {
// 	MustExec(query string, args ...interface{}) sql.Result
// 	Get(dest interface{}, query string, args ...interface{}) error
// }

type DBApi interface {
	Exec(query string, args ...any) (sql.Result, error)
	// Get(dest interface{}, query string, args ...interface{}) error
	// Close...
}
