package database

import (
	"kosyncsrv/types"
)

type DBService interface {
	InitDatabase(schemaUser, schemaDocument string) error
}

type sqlxDBService struct {
	db types.DBApi
}

func NewDBService(
	db types.DBApi,
) DBService {
	return &sqlxDBService{
		db: db,
	}
}

func (s *sqlxDBService) InitDatabase(schemaUser, schemaDocument string) error {
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
