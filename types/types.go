package types

import "database/sql"

type User struct {
	Username string
	Password string
}

type DocumentPosition struct {
	DocumentID string
	Percentage float64
	Progress   string
	Device     string
	DeviceID   string
}

// type DocumentPosition interface {
// 	GetDocumentID() string
// 	GetPercentage() float64
// 	GetProgress()   string
// 	GetDevice()     string
// 	GetDeviceID()   string
// }

type DbUser struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type DbDocumentPosition struct {
	Username   string  `db:"username"`
	DocumentID string  `db:"documentid"`
	Percentage float64 `db:"percentage"`
	Progress   string  `db:"progress"`
	Device     string  `db:"device"`
	DeviceID   string  `db:"device_id"`
	Timestamp  int64   `db:"timestamp"`
}

// func (d *DbDocumentPosition) GetDocumentID() string {
// 	return d.DocumentID
// }

// func (d *DbDocumentPosition) GetPercentage() float64 {
// 	return d.Percentage
// }

// func (d *DbDocumentPosition) GetProgress() string {
// 	return d.Progress
// }

// func (d *DbDocumentPosition) GetDevice() string {
// 	return d.Device
// }

// func (d *DbDocumentPosition) GetDeviceID() string {
// 	return d.DeviceID
// }

type RequestDocumentPosition struct {
	DocumentID string  `json:"document"`
	Percentage float64 `json:"percentage"`
	Progress   string  `json:"progress"`
	Device     string  `json:"device"`
	DeviceID   string  `json:"device_id"`
}

type QueryBuilder interface {
	SchemaUser() string
	SchemaDocument() string
	AddUser() string
	GetUser() string
}

type SqlApi interface {
	Begin() (*sql.Tx, error)
}

type Repo interface {
	InitDatabase() error // doesn't belong to repo...
	AddUser(username, password string) error
	GetUser(username string) (*User, error)
	GetDocumentPosition(username, documentId string) (*DocumentPosition, error)
	UpdateDocumentPosition(username string, documentPosition *DocumentPosition) (*int64, error)
}

type AuthReturnCode int

const (
	Allowed AuthReturnCode = iota
	Forbidden
	Unauthorized
)

type AuthorizationService interface {
	RegisterUser(username, password string) (error, *string)
	AuthorizeUser(username, password string) AuthReturnCode
}

type SyncingService interface {
	GetProgress(username, documentId string) (*DocumentPosition, error)
	UpdateProgress(username string, documentPosition *DocumentPosition) (*int64, error)
}

// type RequestUser struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type RequestHeader struct {
// 	AuthUser string `header:"x-auth-user"`
// 	AuthKey  string `header:"x-auth-key"`
// }
