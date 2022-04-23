package repo_test

import (
	"database/sql"
	"errors"
	"kosyncsrv/database"
	"kosyncsrv/repo"
	"kosyncsrv/test/mocks"

	// test_types "kosyncsrv/test/types"

	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_DBService_InitDatabase_Begin_Fails(t *testing.T) {
	// GIVEN
	userSchema := database.NewQueryBuilder().SchemaUser()
	docSchema := database.NewQueryBuilder().SchemaDocument()

	var tx *sql.Tx
	mockSqlApi := new(mocks.MockedSql)
	mockSqlApi.On("Begin").Return(tx, errors.New("Could not begin transaction"))

	repo := repo.NewRepo(mockSqlApi)
	
	// WHEN
	err := repo.InitDatabase(userSchema, docSchema)

	// THEN
	mockSqlApi.AssertExpectations(t)

	assert.EqualError(t, err, "Could not begin transaction")
}

func Test_DBService_InitDatabase(t *testing.T) {
	// GIVEN
	userSchema := database.NewQueryBuilder().SchemaUser()
	docSchema := database.NewQueryBuilder().SchemaDocument()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// WHEN/THEN
	mock.ExpectBegin()
	ep := mock.ExpectPrepare(userSchema)
	ep.ExpectExec().WillReturnResult(sqlmock.NewResult(1,1))
	ep = mock.ExpectPrepare(docSchema)
	ep.ExpectExec().WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	repo := repo.NewRepo(db)
	err = repo.InitDatabase(userSchema, docSchema)

	assert.Nil(t, err)
}
