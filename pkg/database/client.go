package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/lifina/go-template/pkg/database/mysql"
)

func NewClient(url string) (*sqlx.DB, error) {
	return mysql.NewMySQLClient(url)
}
