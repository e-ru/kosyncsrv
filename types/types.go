package types

import "database/sql"

type SqlApi interface {
	Begin() (*sql.Tx, error)
}

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type Repo interface {
	InitDatabase(schemaUser, schemaDocument string) error
	AddUser(username, password string) bool
	GetUser(username string) (*User, bool)
}

type AuthReturnCode int

const (
	Allowed AuthReturnCode = iota
	Forbidden
	Unauthorized
)

type AuthorizationService interface {
	RegisterUser(username, password string) (bool, string)
	AuthorizeUser(username, password string) (AuthReturnCode, string)
}

type SyncingService interface {
	GetProgress()
	UpdateProgress()
}

// type RequestUser struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type RequestHeader struct {
// 	AuthUser string `header:"x-auth-user"`
// 	AuthKey  string `header:"x-auth-key"`
// }
