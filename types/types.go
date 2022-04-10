package types

import "database/sql"

type DBApi interface {
	MustExec(query string) sql.Result
}
