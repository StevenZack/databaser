package util

import (
	"errors"

	"github.com/StevenZack/databaser/data"
	_ "github.com/go-sql-driver/mysql"
)

func QueryMysql(conn data.Connection, query string) (*data.Result, error) {
	if conn.Type != data.TypeMySQL {
		return nil, errors.New("connection type is not mysql")
	}
	return doSqlQuery(conn, query)
}
