package database

import (
	"database/sql"
	"kosyncsrv/types"
	// _ "github.com/mattn/go-sqlite3"
)

type DBService interface {
	InitDatabase(schemaUser, schemaDocument string) error
}

type sqlDBService struct {
	// db *sql.DB
	dbClient types.SqlApi
}

func NewDBService(
	// db *sql.DB,
	dbClient types.SqlApi,
) DBService {
	return &sqlDBService{
		// db: db,
		dbClient: dbClient,
	}
}

func execStatement(tx *sql.Tx, cmd string) error {
	stmtUser, err := tx.Prepare(cmd)
	if err != nil {
		return err
	}
	defer stmtUser.Close()
	_, err = stmtUser.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlDBService) InitDatabase(schemaUser, schemaDocument string) error {
	tx, err := s.dbClient.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	execStatement(tx, schemaUser)
	execStatement(tx, schemaDocument)

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
