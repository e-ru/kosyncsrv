package repo // whitebox testing

import (
	"database/sql"
	"errors"
	"kosyncsrv/database"

	// "kosyncsrv/repo"
	"kosyncsrv/types"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_DBService_InitDatabase(t *testing.T) {
	// GIVEN
	userSchema := database.NewQueryBuilder().SchemaUser()
	docSchema := database.NewQueryBuilder().SchemaDocument()

	testcases := []struct {
		name       string
		userSchema *string
		docSchema  *string
		err        error
		wantErr    bool
	}{
		{name: "init db successfully", userSchema: &userSchema, docSchema: &docSchema, wantErr: false},
		{name: "init db unsuccessfully", docSchema: &docSchema, err: errors.New("Could not create user table"), wantErr: true},
		{name: "init db unsuccessfully", userSchema: &userSchema, err: errors.New("Could not create document table"), wantErr: true},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if testcase.wantErr {
				if testcase.userSchema == nil {
					mock.ExpectExec(database.NewQueryBuilder().SchemaUser()).WillReturnError(testcase.err)
				}
				if testcase.docSchema == nil {
					mock.ExpectExec(userSchema).WillReturnResult(sqlmock.NewResult(1, 1))
					mock.ExpectExec(database.NewQueryBuilder().SchemaDocument()).WillReturnError(testcase.err)
				}
			} else {
				mock.ExpectExec(userSchema).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec(docSchema).WillReturnResult(sqlmock.NewResult(1, 1))
			}

			// WHEN
			repo := NewRepo(db, database.NewQueryBuilder())
			err = repo.InitDatabase()

			// THEN
			if testcase.wantErr {
				assert.Error(t, testcase.err)
			} else {
				assert.Nil(t, err)
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
		{name: "get user unsuccessfully no such user", query: database.NewQueryBuilder().GetUser(), err: sql.ErrNoRows, wantErr: true},
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
			repo := NewRepo(db, database.NewQueryBuilder())

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

func Test_DBService_Add_User(t *testing.T) {
	// GIVEN
	username := "username"
	password := "password"

	testcases := []struct {
		name       string
		getUserErr error
		query      string
		err        error
		wantErr    bool
	}{
		{name: "add user successfully", getUserErr: sql.ErrNoRows, query: database.NewQueryBuilder().AddUser(), wantErr: false},
		{name: "add user unsuccessfully", getUserErr: sql.ErrNoRows, query: database.NewQueryBuilder().AddUser(), err: errors.New("Exec Error"), wantErr: true},
		{name: "add user unsuccessfully", getUserErr: sql.ErrConnDone, query: database.NewQueryBuilder().AddUser(), err: errors.New("Could not add user"), wantErr: true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if testcase.wantErr {
				if testcase.getUserErr == sql.ErrNoRows {
					mock.ExpectExec(testcase.query).WithArgs(username, password).WillReturnError(testcase.err)
				}
			} else {
				mock.ExpectExec(testcase.query).WithArgs(username, password).WillReturnResult(sqlmock.NewResult(1, 1))
			}

			repo := NewRepo(db, database.NewQueryBuilder())

			called := false
			getUser = func(s *sqlRepo, username string) (*types.User, error) {
				called = true
				return nil, testcase.getUserErr
			}

			// WHEN
			err = repo.AddUser(username, password)

			// THEN
			if !called {
				t.Errorf("GetUser was not called")
			}

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

func Test_DBService_Get_Document_Position_By_User_Id(t *testing.T) {
	// GIVEN
	username := "username"
	documentId := "documentId"

	testcases := []struct {
		name    string
		query   string
		ret     *types.DocumentPosition
		err     error
		wantErr bool
	}{
		{
			name:  "get document position successfully",
			query: database.NewQueryBuilder().GetDocumentPositionByUserId(),
			ret: &types.DocumentPosition{
				Username:   username,
				DocumentID: documentId,
				Percentage: 5.5,
				Progress:   "5",
				Device:     "Dev",
				DeviceID:   "DevId",
				Timestamp:  5,
			},
			wantErr: false,
		},
		{
			name: "get document position unsuccessfully", 
			query: database.NewQueryBuilder().GetDocumentPositionByUserId(), 
			err: errors.New("Exec Error"), 
			wantErr: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if testcase.wantErr {
				mock.ExpectQuery(testcase.query).WithArgs(documentId, username).WillReturnError(testcase.err)
			} else {
				newRows := sqlmock.NewRows(
					[]string{
						username,
						documentId,
						"percentage",
						"progress",
						"device",
						"device_id",
						"timestamp",
					}).AddRow(
					username,
					documentId,
					5.5,
					"5",
					"Dev",
					"DevId",
					5,
				)
				mock.ExpectQuery(testcase.query).WithArgs(documentId, username).WillReturnRows(newRows)
			}

			repo := NewRepo(db, database.NewQueryBuilder())

			// WHEN
			docPos, err := repo.GetDocumentPositionByUserId(documentId, username)

			if testcase.wantErr {
				if err == nil {
					t.Errorf("was expecting an error, but there was none")
				}
			} else {
				if err != nil {
					t.Errorf("error was not expected while updating stats: %s", err)
				}
				assert.Equal(t, testcase.ret, docPos)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_DBService_Get_Document_Position_By_Device_Id(t *testing.T) {
	// GIVEN
	username := "username"
	documentId := "documentId"
	deviceId := "deviceId"

	testcases := []struct {
		name    string
		query   string
		ret     *types.DocumentPosition
		err     error
		wantErr bool
	}{
		{
			name:  "get document position successfully",
			query: database.NewQueryBuilder().GetDocumentPositionByDeviceId(),
			ret: &types.DocumentPosition{
				Username:   username,
				DocumentID: documentId,
				Percentage: 5.5,
				Progress:   "5",
				Device:     "Dev",
				DeviceID:   "DevId",
				Timestamp:  5,
			},
			wantErr: false,
		},
		{
			name: "get document position unsuccessfully", 
			query: database.NewQueryBuilder().GetDocumentPositionByDeviceId(), 
			err: errors.New("Exec Error"), 
			wantErr: true,
		},
		{
			name: "get document position unsuccessfully no such document", 
			query: database.NewQueryBuilder().GetDocumentPositionByDeviceId(), 
			err: sql.ErrNoRows, 
			wantErr: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if testcase.wantErr {
				mock.ExpectQuery(testcase.query).WithArgs(documentId, deviceId).WillReturnError(testcase.err)
			} else {
				newRows := sqlmock.NewRows(
					[]string{
						username,
						documentId,
						"percentage",
						"progress",
						"device",
						"device_id",
						"timestamp",
					}).AddRow(
					username,
					documentId,
					5.5,
					"5",
					"Dev",
					"DevId",
					5,
				)
				mock.ExpectQuery(testcase.query).WithArgs(documentId, deviceId).WillReturnRows(newRows)
			}

			repo := NewRepo(db, database.NewQueryBuilder())

			// WHEN
			docPos, err := repo.GetDocumentPositionByDeviceId(documentId, deviceId)

			if testcase.wantErr {
				if err == nil {
					t.Errorf("was expecting an error, but there was none")
				}
			} else {
				if err != nil {
					t.Errorf("error was not expected while updating stats: %s", err)
				}
				assert.Equal(t, testcase.ret, docPos)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_DBService_Update_Document_Position(t *testing.T) {
	// GIVEN
	username := "username"
	docPos := types.DocumentPosition{
		Username:   username,
		DocumentID: "documentId",
		Percentage: 5.5,
		Progress:   "5",
		Device:     "Dev",
		DeviceID:   "DevId",
		Timestamp:  5,
	}

	testcases := []struct {
		name       string
		docExists  bool
		docExistsQuery      string
		err        error
		wantErr    bool
	}{
		// {name: "add user successfully", docExists: true, docExistsQuery: database.NewQueryBuilder().DocumentExists(), wantErr: false},
		// {name: "add user unsuccessfully", getUserErr: sql.ErrNoRows, query: database.NewQueryBuilder().AddUser(), err: errors.New("Exec Error"), wantErr: true},
		// {name: "add user unsuccessfully", getUserErr: sql.ErrConnDone, query: database.NewQueryBuilder().AddUser(), err: errors.New("Could not add user"), wantErr: true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if testcase.wantErr {
				mock.ExpectExec(testcase.docExistsQuery).WithArgs(docPos.DocumentID, docPos.DeviceID).WillReturnError(testcase.err)
			} else {
				mock.ExpectExec(testcase.docExistsQuery).WithArgs(docPos.DocumentID, docPos.DeviceID).WillReturnResult(sqlmock.NewResult(1, 1))
			}

			// repo := NewRepo(db, database.NewQueryBuilder())

			// WHEN

		})
	}
}
