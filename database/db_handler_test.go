package database_test

import (
	"errors"
	"kosyncsrv/database"
	"kosyncsrv/test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testResult struct {
}

func (dr testResult) LastInsertId() (int64, error) {
	return 0, errors.New("foo")
}

func (dr testResult) RowsAffected() (int64, error) {
	return 0, errors.New("foo")
}

func Test_DBHandler_InitDatabase_Panics(t *testing.T) {
	// GIVEN
	query1 := "SELECT *"
	query := "SELECT *  "

	t.Run("InitDatabase panics with wrong init statements", func(t *testing.T) {
		mockDb := new(mocks.MockedDB)

		mockDb.On("MustExec", query1, mock.Anything).Return(testResult{})
		mockDb.On("MustExec", query, mock.Anything).Return(testResult{})

		dbHandler := database.NewDBHandler(mockDb)

		// WHEN/THEN
		assert.Panics(t, func() { dbHandler.InitDatabase() }, "The code did not panic")

	})
}

func Test_sqlxDBHandler_InitDatabase_Doesnt_Panic(t *testing.T) {
	// GIVEN
	query1 := "SELECT *"
	query := "SELECT * "

	t.Run("InitDatabase doesn't panic with correct init statements", func(t *testing.T) {
		mockDb := new(mocks.MockedDB)

		mockDb.On("MustExec", query1, mock.Anything).Return(testResult{})
		mockDb.On("MustExec", query, mock.Anything).Return(testResult{})

		dbHandler := database.NewDBHandler(mockDb)

		// WHEN/THEN
		assert.NotPanics(t, func() { dbHandler.InitDatabase() }, "The code did panic")

	})
}
