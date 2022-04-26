package repo

import (
	"database/sql"
	"kosyncsrv/types"
)

type sqlRepo struct {
	sqlClient    types.SqlApi
	queryBuilder types.QueryBuilder
}

func NewRepo(
	sqlClient types.SqlApi,
	queryBuilder types.QueryBuilder,
) types.Repo {
	return &sqlRepo{
		sqlClient:    sqlClient,
		queryBuilder: queryBuilder,
	}
}

func execStatement(tx *sql.Tx, cmd string, args ...any) error {
	stmt, err := tx.Prepare(cmd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if len(args) > 0 {
		_, err = stmt.Exec(args...)
	} else {
		_, err = stmt.Exec()
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlRepo) InitDatabase() error {
	tx, err := s.sqlClient.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = execStatement(tx, s.queryBuilder.SchemaUser())
	if err != nil {
		return err
	}
	err = execStatement(tx, s.queryBuilder.SchemaDocument())
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
	err = execStatement(tx, s.queryBuilder.AddUser(), username, password)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlRepo) GetUser(username string) (*types.User, error) {
	return nil, nil
}

func (mr *sqlRepo) GetDocumentPosition(username, documentId string) (*types.DocumentPosition, error) {
	return nil, nil
}

func (mr *sqlRepo) UpdateDocumentPosition(username string, documentPosition *types.DocumentPosition) (*int64, error) {
	timestamp := int64(0)
	return &timestamp, nil
}
