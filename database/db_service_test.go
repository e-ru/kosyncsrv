package database_test

import (
	"database/sql"
	"errors"
	"kosyncsrv/database"
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
	mockDbApi := new(mocks.MockedDB)
	mockDbApi.On("Begin").Return(tx, errors.New("Could not begin transaction"))

	dbService := database.NewDBService(mockDbApi)
	
	// WHEN
	err := dbService.InitDatabase(userSchema, docSchema)

	// THEN
	mockDbApi.AssertExpectations(t)

	assert.EqualError(t, err, "Could not begin transaction")
}

func Test_DBService_InitDatabase_Commit_Is_Called(t *testing.T) {
	// GIVEN
	userSchema := database.NewQueryBuilder().SchemaUser()
	docSchema := database.NewQueryBuilder().SchemaDocument()

	// var tx *sql.Tx
	mockDbApi := new(mocks.MockedDB)
	mockTx := new(mocks.MockSqlTx)
	mockDbApi.On("Begin").Return(mockTx, nil)
	mockTx.On("Commit").Return(nil)

	dbService := database.NewDBService(mockDbApi)
	
	// WHEN
	err := dbService.InitDatabase(userSchema, docSchema)

	// THEN
	mockDbApi.AssertExpectations(t)
	mockTx.AssertExpectations(t)

	assert.NoError(t, err)
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

	dbService := database.NewDBService(db)
	err = dbService.InitDatabase(userSchema, docSchema)

	assert.Nil(t, err)
}
