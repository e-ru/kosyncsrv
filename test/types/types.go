package test_types

type TestResult struct {
}

func (dr TestResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (dr TestResult) RowsAffected() (int64, error) {
	return 0, nil
}