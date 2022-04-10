package database

type QueryBuilder interface {
	SchemaUser() string
	SchemaDocument() string
}

type sqlite3QueryBuilder struct {
}

func NewQueryBuilder() QueryBuilder {
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
