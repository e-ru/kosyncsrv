package mocks

import (
	// "database/sql"
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockedDB struct {
	mock.Mock
}

func (m *MockedDB) Begin() (*sql.Tx, error) {
	args := m.Called()
	if args.Get(0).(*sql.Tx) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*sql.Tx), nil
	}
}

type MockSqlTx struct {
	*sql.Tx
	mock.Mock
}

func (t *MockSqlTx) Commit() error {
	args := t.Called()
	return args.Error(0)
}
