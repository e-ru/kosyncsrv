package database

import "kosyncsrv/types"

type sqlite3QueryBuilder struct {
}

func NewQueryBuilder() types.QueryBuilder {
	return &sqlite3QueryBuilder{}
}

func (q *sqlite3QueryBuilder) SchemaUser() string {
	return `
CREATE TABLE IF NOT EXISTS "user" (
	"username"  TEXT(255),
	"password"  TEXT(255)
);
`
}

func (q *sqlite3QueryBuilder) SchemaDocument() string {
	return `
CREATE TABLE IF NOT EXISTS "document" (
	"username"  TEXT(255),
	"documentid"  TEXT(255),
	"percentage"  REAL(64,4),
	"progress"  TEXT(255),
	"device"  TEXT(255),
	"device_id"  TEXT(255),
	"timestamp"  INTEGER
);
`
}

func (q *sqlite3QueryBuilder) AddUser() string {
	return `INSERT INTO user (username, password) VALUES ($1, $2)`
}

func (q *sqlite3QueryBuilder) GetUser() string {
	return `SELECT * FROM user WHERE username=$1`
}

func (q *sqlite3QueryBuilder) DocumentExists() string {
	return `SELECT * FROM document WHERE documentid=$1 AND device_id=$2`
}

func (q *sqlite3QueryBuilder) GetDocumentPosition() string {
	return "SELECT * FROM document WHERE document.username=$1 AND document.documentid=$2 ORDER BY document.timestamp DESC"
}
