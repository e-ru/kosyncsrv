package mocks

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockedDB struct {
	mock.Mock
}

func (m *MockedDB) MustExec(query string, args ...interface{}) sql.Result {
	calledArgs := m.Called(query, args)
	return calledArgs.Get(0).(sql.Result)
}