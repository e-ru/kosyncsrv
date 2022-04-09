package database

import (
	"database/sql"
	"fmt"
	"kosyncsrv/types"
)

type DBHandler interface {
	InitDatabase() sql.Result
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

func (s *sqlxDBHandler) InitDatabase() sql.Result {
	res :=  s.db.MustExec("SELECT *")
	s.db.MustExec("SELECT * ")

	fmt.Printf("res: %+v", res)
	// s.db.MustExec(schemaUser)
	// s.db.MustExec(schemaDocument)

	return res
}
