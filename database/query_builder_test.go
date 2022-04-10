package database_test

import (
	"kosyncsrv/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueryBuilder_UserSchema(t *testing.T) {
	// GIVEN
	expSchemaUser := `
CREATE TABLE IF NOT EXISTS "user" (
	"username"  TEXT(255),
	"password"  TEXT(255)
);
`

	t.Run("userSchema should return correct create statement", func(t *testing.T) {
		queryBuilder := database.NewQueryBuilder()

		// WHEN
		schemaUser := queryBuilder.SchemaUser()

		//THEN
		assert.Equal(t,expSchemaUser, schemaUser)
	})
}