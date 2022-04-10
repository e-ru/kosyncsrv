package database_test

import (
	"kosyncsrv/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueryBuilder_UserSchema(t *testing.T) {
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
