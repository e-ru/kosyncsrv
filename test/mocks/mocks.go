package mocks

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
	// "github.com/DATA-DOG/go-sqlmock"
	// "github.com/jmoiron/sqlx"

)

type MockedDB struct {
	// *sql.DB
	mock.Mock
}

func (m *MockedDB) MustExec(query string, args ...interface{}) sql.Result {
	calledArgs := m.Called(query)
	return calledArgs.Get(0).(sql.Result)
}

// func (m *MockedDB) MustExec(query string, args ...interface{}) sql.Result {
// 	calledArgs := m.Called(query)
// 	return calledArgs.Get(0).(sql.Result)
// }

func (m *MockedDB) Get(dest interface{}, query string, args ...interface{}) error {
	return nil
}