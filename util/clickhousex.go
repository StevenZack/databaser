package util

import (
	"errors"

	_ "github.com/ClickHouse/clickhouse-go"

	"github.com/StevenZack/databaser/data"
)

func QueryClickhouse(conn data.Connection, query string) (*data.Result, error) {
	if conn.Type != data.TypeClickhouse {
		return nil, errors.New("connection type is not clickhouse")
	}
	return doSqlQuery(conn, query)
}
