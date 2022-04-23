package repo

import (
	"database/sql"
	"kosyncsrv/types"
)

type sqlRepo struct {
	sqlClient types.SqlApi
}

func NewRepo(sqlClient types.SqlApi) types.Repo {
	return &sqlRepo{sqlClient: sqlClient}
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

func (s *sqlRepo) InitDatabase(schemaUser, schemaDocument string) error {
	tx, err := s.sqlClient.Begin()
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

func (s *sqlRepo) AddUser(username, password string) bool {
	return false
}
