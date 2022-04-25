package repo_test

import (
	"database/sql"
	"errors"
	"kosyncsrv/database"
	"kosyncsrv/repo"
	"kosyncsrv/test/mocks"

	"testing"

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

	mock.ExpectBegin()
	ep := mock.ExpectPrepare(userSchema)
	ep.ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
	ep = mock.ExpectPrepare(docSchema)
	ep.ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// WHEN
	repo := repo.NewRepo(db)
	err = repo.InitDatabase(userSchema, docSchema)

	// THEN
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.Nil(t, err)
}

func Test_DBService_Add_User(t *testing.T) {
	// GIVEN
	username := "username"
	password := "password"

	testcases := []struct {
		name    string
		query   string
		err     error
		wantErr bool
	}{
		{name: "add user successfully", query: database.NewQueryBuilder().AddUser(), wantErr: false},
		{name: "add user unsuccessfully", query: database.NewQueryBuilder().AddUser(), err: errors.New("Could not add user"), wantErr: true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectBegin()
			ep := mock.ExpectPrepare(testcase.query)
			if testcase.wantErr {
				ep.ExpectExec().WithArgs(username, password).WillReturnError(testcase.err)
				mock.ExpectRollback()
			} else {
				ep.ExpectExec().WithArgs(username, password).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			}

			// WHEN
			repo := repo.NewRepo(db)
			if testcase.wantErr {
				if err := repo.AddUser(username, password); err == nil {
					t.Errorf("was expecting an error, but there was none")
				}
			} else {			
				if err := repo.AddUser(username, password); err != nil {
					t.Errorf("error was not expected while updating stats: %s", err)
				}
			}

			// THEN
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			assert.NoError(t, err)
		})
	}
}
