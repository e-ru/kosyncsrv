package database

import (
	"kosyncsrv/types"
)

type DBHandler interface {
	InitDatabase() error
}

type sqlxDBHandler struct {
	db types.DBApi
}

func NewDBHandler(
	db types.DBApi,
) DBHandler {
	return &sqlxDBHandler{
		db: db,
	}
}

func (s *sqlxDBHandler) InitDatabase() error {
	res := s.db.MustExec("SELECT *")
	_, err := res.LastInsertId()
	if err != nil {
		return err
	}

	res = s.db.MustExec("SELECT * ")
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
