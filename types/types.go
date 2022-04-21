package types

import "database/sql"

type SqlApi interface {
	Begin() (*sql.Tx, error)
}

type SqlTxApi interface {
	Commit() error
}

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
