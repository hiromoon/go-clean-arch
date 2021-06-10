package infra

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
}

const dataSourceName = "root:root@tcp(127.0.0.1:3306)/test?parseTime=true"

func ConnectDatabase() (*Database, error) {
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
