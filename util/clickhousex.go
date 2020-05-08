package util

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"

	"github.com/StevenZack/databaser/data"
)

func QueryClickhouse(conn data.Connection, query string) (*data.Result, error) {
	result := &data.Result{
		Connection: conn.Name,
		Query:      query,
	}
	if conn.Type != data.TypeClickhouse {
		return result, errors.New("connection type is not clickhouse")
	}
	db, e := sql.Open("clickhouse", conn.Dsn)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	defer db.Close()
	rows, e := db.Query(query)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	defer rows.Close()

	for rows.Next() {
		if len(result.Rows) == 500 {
			break
		}
		result.Columns, e = rows.Columns()
		if e != nil {
			log.Println(e)
			return nil, e
		}
		vs := make([]interface{}, len(result.Columns))
		uvs := make([]interface{}, len(result.Columns))
		for i := range vs {
			uvs[i] = &vs[i]
		}
		e = rows.Scan(uvs...)
		if e != nil {
			log.Println(e)
			return nil, e
		}
		row := data.Row{}
		for _, v := range vs {
			row.Values = append(row.Values, fmt.Sprint(v))
		}
		result.Rows = append(result.Rows, row)
	}

	return result, nil
}
