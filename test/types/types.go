package test_types

import "errors"


type TestResult struct {
}

func (dr TestResult) LastInsertId() (int64, error) {
	return 0, errors.New("foo")
}

func (dr TestResult) RowsAffected() (int64, error) {
	return 0, errors.New("foo")
}