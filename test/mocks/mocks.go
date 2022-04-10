package mocks

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockedDB struct {
	mock.Mock
}

func (m *MockedDB) MustExec(query string) sql.Result {
	args := m.Called(query)
	return args.Get(0).(sql.Result)
}
