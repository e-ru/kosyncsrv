package database

import (
	"kosyncsrv/types"
)

type DBHandler interface {
	InitDatabase(schemaUser, schemaDocument string) error
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

func (s *sqlxDBHandler) InitDatabase(schemaUser, schemaDocument string) error {
	res := s.db.MustExec(schemaUser)
	_, err := res.LastInsertId()
	if err != nil {
		return err
	}

	res = s.db.MustExec(schemaDocument)
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
