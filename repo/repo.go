package repo

import (
	"database/sql"
	"fmt"
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

func (s *sqlRepo) InitDatabase() error {
	if _, err := s.sqlClient.Exec(s.queryBuilder.SchemaUser()); err != nil {
		return err
	}
	if _, err := s.sqlClient.Exec(s.queryBuilder.SchemaDocument()); err != nil {
		return err
	}
	return nil
}

func (s *sqlRepo) AddUser(username, password string) error {
	if _, err := s.sqlClient.Exec(s.queryBuilder.AddUser(), username, password); err != nil {
		return err
	}
	return nil
}

func (s *sqlRepo) GetUser(username string) (*types.User, error) {
	var user types.User

	row := s.sqlClient.QueryRow(s.queryBuilder.GetUser(), username)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No such User: %+v", username)
		}
		return nil, err
	}
	return &user, nil
}

func (mr *sqlRepo) GetDocumentPosition(username, documentId string) (*types.DocumentPosition, error) {
	return nil, nil
}

func (mr *sqlRepo) UpdateDocumentPosition(username string, documentPosition *types.DocumentPosition) (*int64, error) {
	timestamp := int64(0)
	return &timestamp, nil
}
