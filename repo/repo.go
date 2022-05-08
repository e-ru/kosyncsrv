package repo

import (
	"database/sql"
	"fmt"
	"kosyncsrv/types"
	"log"
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

func (s *sqlRepo) GetUser(username string) (*types.User, error) {
	log.Printf("Get user: %+v", username)

	var user types.User
	row := s.sqlClient.QueryRow(s.queryBuilder.GetUser(), username)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		log.Printf("Get user error: %+v", err)
		return nil, err
	}
	return &user, nil
}

var getUser = (*sqlRepo).GetUser

func (s *sqlRepo) AddUser(username, password string) error {
	log.Printf("Add user: %+v", username)

	_, err := getUser(s, username)
	if err == sql.ErrNoRows {
		if _, err := s.sqlClient.Exec(s.queryBuilder.AddUser(), username, password); err != nil {
			return fmt.Errorf("Could not add User, error: %+v", err)
		}
		return nil
	}
	return fmt.Errorf("Could not add User, error: %+v", err)
}

func (mr *sqlRepo) GetDocumentPosition(username, documentId string) (*types.DocumentPosition, error) {
	return nil, nil
}

func (mr *sqlRepo) UpdateDocumentPosition(username string, documentPosition *types.DocumentPosition) (*int64, error) {
	timestamp := int64(0)
	return &timestamp, nil
}
