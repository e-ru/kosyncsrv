package database_test

import (
	"kosyncsrv/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueryBuilder(t *testing.T) {
	// GIVEN

	queryBuilder := database.NewQueryBuilder()
	testCases := []struct {
		name   string
		expSql string
		query  func() string
	}{
		{
			name: "userSchema should return correct create statement",
			expSql: `
CREATE TABLE IF NOT EXISTS "user" (
	"username"  TEXT(255),
	"password"  TEXT(255)
);
`,
			query: queryBuilder.SchemaUser,
		},
		{
			name: "DocumentSchema should return correct create statement",
			expSql: `
CREATE TABLE IF NOT EXISTS "document" (
	"username"  TEXT(255),
	"documentid"  TEXT(255),
	"percentage"  REAL(64,4),
	"progress"  TEXT(255),
	"device"  TEXT(255),
	"device_id"  TEXT(255),
	"timestamp"  INTEGER
);
`,
			query: queryBuilder.SchemaDocument,
		},
		{
			name:   "AddUser should return correct statement",
			expSql: "INSERT INTO user (username, password) VALUES ($1, $2)",
			query:  queryBuilder.AddUser,
		},
		{
			name:   "GetUser should return correct statement",
			expSql: "SELECT * FROM user WHERE username=$1",
			query:  queryBuilder.GetUser,
		},
		{
			name:   "DocumentExists should return correct statement",
			expSql: "SELECT * FROM document WHERE documentid=$1 AND device_id=$2",
			query:  queryBuilder.DocumentExists,
		},
		{
			name:   "DocumentExists should return correct statement",
			expSql: "SELECT * FROM document WHERE document.username=$1 AND document.documentid=$2 ORDER BY document.timestamp DESC",
			query:  queryBuilder.GetDocumentPosition,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			// WHEN
			sql := testCase.query()

			//THEN
			assert.Equal(t, testCase.expSql, sql)
		})
	}

}
