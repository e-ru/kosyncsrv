package repo_test

import (
	"errors"
	"kosyncsrv/database"
	"kosyncsrv/repo"
	"kosyncsrv/types"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_DBService_InitDatabase(t *testing.T) {
	// GIVEN
	userSchema := database.NewQueryBuilder().SchemaUser()
	docSchema := database.NewQueryBuilder().SchemaDocument()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(userSchema).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(docSchema).WillReturnResult(sqlmock.NewResult(1, 1))

	// WHEN
	repo := repo.NewRepo(db, database.NewQueryBuilder())
	err = repo.InitDatabase()

	// THEN
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
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

			if testcase.wantErr {
				mock.ExpectExec(testcase.query).WithArgs(username, password).WillReturnError(testcase.err)
			} else {
				mock.ExpectExec(testcase.query).WithArgs(username, password).WillReturnResult(sqlmock.NewResult(1, 1))
			}

			repo := repo.NewRepo(db, database.NewQueryBuilder())

			// WHEN
			err = repo.AddUser(username, password)

			// THEN
			if testcase.wantErr {
				if err == nil {
					t.Errorf("was expecting an error, but there was none")
				}
			} else {
				if err != nil {
					t.Errorf("error was not expected while updating stats: %s", err)
				}
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_DBService_Get_User(t *testing.T) {
	// GIVEN
	username := "username"
	password := "password"

	testcases := []struct {
		name    string
		query   string
		user    *types.User
		err     error
		wantErr bool
	}{
		{name: "get user successfully", query: database.NewQueryBuilder().GetUser(), user: &types.User{Username: username, Password: password}, wantErr: false},
		{name: "get user unsuccessfully", query: database.NewQueryBuilder().GetUser(), err: errors.New("Could not get user"), wantErr: true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if testcase.wantErr {
				mock.ExpectQuery(testcase.query).WithArgs(username).WillReturnError(testcase.err)
			} else {
				newRows := sqlmock.NewRows([]string{username, password}).AddRow(username, password)
				mock.ExpectQuery(testcase.query).WithArgs(username).WillReturnRows(newRows)
			}
			repo := repo.NewRepo(db, database.NewQueryBuilder())

			// WHEN
			user, err := repo.GetUser(username)

			// THEN
			if testcase.wantErr {
				if err == nil {
					t.Errorf("was expecting an error, but there was none")
				}
			} else {
				if err != nil {
					t.Errorf("error was not expected while updating stats: %s", err)
				}
				assert.Equal(t, testcase.user, user)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
