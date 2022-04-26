package mocks

import (
	"database/sql"
	"kosyncsrv/types"

	"github.com/stretchr/testify/mock"
)

type MockedSql struct {
	mock.Mock
}

func (ms *MockedSql) Begin() (*sql.Tx, error) {
	args := ms.Called()
	if args.Get(0).(*sql.Tx) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*sql.Tx), nil
	}
}

type MockedRepo struct {
	mock.Mock
}

func (mr *MockedRepo) InitDatabase() error {
	args := mr.Called()
	return args.Error(0)
}

func (mr *MockedRepo) AddUser(username, password string) error {
	args := mr.Called(username, password)
	return args.Error(0)
}

func (mr *MockedRepo) GetUser(username string) (*types.User, error) {
	args := mr.Called(username)
	return args.Get(0).(*types.User), args.Error(1)
}

func (mr *MockedRepo) GetDocumentPosition(username, documentId string) (*types.DocumentPosition, error) {
	args := mr.Called(username, documentId)
	return args.Get(0).(*types.DocumentPosition), args.Error(1)
}

func (mr *MockedRepo) UpdateDocumentPosition(username string, documentPosition *types.DocumentPosition) (*int64, error) {
	args := mr.Called(username, documentPosition)
	return args.Get(0).(*int64), args.Error(1)
}

// func (mr *MockedRepo) AuthorizeUser(username, password string) (types.AuthReturnCode, string) {
// 	args := mr.Called(username, password)
// 	return args.Get(0).(types.AuthReturnCode), args.Get(1).(string)
// }
