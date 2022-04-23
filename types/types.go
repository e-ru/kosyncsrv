package types

import "database/sql"

type SqlApi interface {
	Begin() (*sql.Tx, error)
}

type Repo interface {
	InitDatabase(schemaUser, schemaDocument string) error
	AddUser(username, password string) bool
}

type SyncingService interface {
	GetProgress()
	UpdateProgress()
}

type AuthorizationService interface {
	RegisterUser(username, password string) (bool, string)
	AuthorizeUser()
}

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
