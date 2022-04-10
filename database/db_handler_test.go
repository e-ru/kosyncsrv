package database_test

import (
	"kosyncsrv/database"
	"kosyncsrv/test/mocks"
	test_types "kosyncsrv/test/types"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DBHandler_InitDatabase(t *testing.T) {
	// GIVEN

	testCases := []struct {
		name                 string
		wantedSchemaUser     string
		wantedSchemaDocument string
		passedSchemaUser     string
		passedSchemaDocument string
		panics               bool
	}{
		{
			name:                 "InitDatabase panics with wrong init statements",
			wantedSchemaUser:     "SELECT *",
			wantedSchemaDocument: "SELECT *",
			passedSchemaUser:     "SELECT *",
			passedSchemaDocument: "SELECT * ",
			panics:               true,
		},
		{
			name:                 "InitDatabase doesn't panic with correct init statements",
			wantedSchemaUser:     "SELECT *",
			wantedSchemaDocument: "SELECT *",
			passedSchemaUser:     "SELECT *",
			passedSchemaDocument: "SELECT *",
			panics:               false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockDb := new(mocks.MockedDB)

			mockDb.On("MustExec", testCase.wantedSchemaUser).Return(test_types.TestResult{})
			mockDb.On("MustExec", testCase.wantedSchemaDocument).Return(test_types.TestResult{})

			dbHandler := database.NewDBHandler(mockDb)

			// WHEN/THEN
			if testCase.panics {
				assert.Panics(
					t,
					func() { dbHandler.InitDatabase(testCase.passedSchemaUser, testCase.passedSchemaDocument) },
					"The code did not panic",
				)
			} else {
				assert.NotPanics(
					t,
					func() { dbHandler.InitDatabase(testCase.passedSchemaUser, testCase.passedSchemaDocument) },
					"The code did panic",
				)
				mockDb.AssertExpectations(t)
			}
		})
	}

}
