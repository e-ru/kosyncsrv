package database_test

import (
	"kosyncsrv/database"
	// test_types "kosyncsrv/test/types"

	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// func Test_DBService_InitDatabase(t *testing.T) {
// 	// GIVEN

// 	testCases := []struct {
// 		name                 string
// 		wantedSchemaUser     string
// 		wantedSchemaDocument string
// 		passedSchemaUser     string
// 		passedSchemaDocument string
// 		panics               bool
// 	}{
// 		{
// 			name:                 "InitDatabase panics with wrong init statements",
// 			wantedSchemaUser:     "SELECT *",
// 			wantedSchemaDocument: "SELECT *",
// 			passedSchemaUser:     "SELECT *",
// 			passedSchemaDocument: "SELECT * ",
// 			panics:               true,
// 		},
// 		{
// 			name:                 "InitDatabase doesn't panic with correct init statements",
// 			wantedSchemaUser:     "SELECT *",
// 			wantedSchemaDocument: "SELECT *",
// 			passedSchemaUser:     "SELECT *",
// 			passedSchemaDocument: "SELECT *",
// 			panics:               false,
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			mockDb := new(mocks.MockedDB)

// 			mockDb.On("MustExec", testCase.wantedSchemaUser).Return(test_types.TestResult{})
// 			mockDb.On("MustExec", testCase.wantedSchemaDocument).Return(test_types.TestResult{})

// 			dbHandler := database.NewDBService(mockDb)

// 			// WHEN/THEN
// 			if testCase.panics {
// 				assert.Panics(
// 					t,
// 					func() { dbHandler.InitDatabase(testCase.passedSchemaUser, testCase.passedSchemaDocument) },
// 					"The code did not panic",
// 				)
// 			} else {
// 				assert.NotPanics(
// 					t,
// 					func() { dbHandler.InitDatabase(testCase.passedSchemaUser, testCase.passedSchemaDocument) },
// 					"The code did panic",
// 				)
// 				mockDb.AssertExpectations(t)
// 			}
// 		})
// 	}

// }

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
	// mock.ExpectExec(userSchema).WillReturnResult(sqlmock.NewResult(1,1))
	// mock.ExpectExec(docSchema).WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	dbService := database.NewDBService(db)
	err = dbService.InitDatabase(userSchema, docSchema)

	assert.Nil(t, err)
}
