package repo

import (
	"database/sql"
	"kosyncsrv/database"
	"kosyncsrv/types"
)

type sqlRepo struct {
	sqlClient types.SqlApi
}

func NewRepo(sqlClient types.SqlApi) types.Repo {
	return &sqlRepo{sqlClient: sqlClient}
}

func execStatement(tx *sql.Tx, cmd string, args ...any) error {
	stmtUser, err := tx.Prepare(cmd)
	if err != nil {
		return err
	}
	defer stmtUser.Close()
	if len(args) > 0 {
		_, err = stmtUser.Exec(args...)
	} else {
		_, err = stmtUser.Exec()
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlRepo) InitDatabase(schemaUser, schemaDocument string) error {
	tx, err := s.sqlClient.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = execStatement(tx, schemaUser)
	if err != nil {
		return err
	}
	err = execStatement(tx, schemaDocument)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlRepo) AddUser(username, password string) error {
	tx, err := s.sqlClient.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = execStatement(tx, database.NewQueryBuilder().AddUser(), username, password)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlRepo) GetUser(username string) (*types.User, bool) {
	return nil, false
}

func (mr *sqlRepo) GetDocumentPosition(username, documentId string) (*types.DocumentPosition, error) {
	return nil, nil
}

func (mr *sqlRepo) UpdateDocumentPosition(username string, documentPosition *types.DocumentPosition) (*int64, error) {
	timestamp := int64(0)
	return &timestamp, nil
}
