package database

import (
	"database/sql"
	// _ "github.com/mattn/go-sqlite3"
)

type DBService interface {
	InitDatabase(schemaUser, schemaDocument string) error
}

type sqlxDBService struct {
	db *sql.DB
}

func NewDBService(
	db *sql.DB,
) DBService {
	return &sqlxDBService{
		db: db,
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

func (s *sqlxDBService) InitDatabase(schemaUser, schemaDocument string) error {
	tx, err := s.db.Begin()
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
	// tx, err := s.db.Begin()
	// if err != nil {
	// 	return err
	// }
	// defer func() {
	// 	switch err {
	// 	case nil:
	// 		err = tx.Commit()
	// 	default:
	// 		tx.Rollback()
	// 	}
	// }()

	// res, err := tx.Exec(schemaUser)
	// fmt.Printf("res: %+v", res)
	// if err != nil {
	// 	return err
	// }

	// return nil
	// if _, err := s.db.Exec(schemaDocument); err != nil {
	// 	panic(err)
	// }

	// res := s.db.MustExec(schemaUser)
	// _, err := res.LastInsertId()
	// if err != nil {
	// 	return err
	// }

	// res = s.db.MustExec(schemaDocument)
	// _, err = res.RowsAffected()
	// if err != nil {
	// 	return err
	// }

	// return nil
}
