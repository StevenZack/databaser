package util

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/StevenZack/databaser/data"
)

func doSqlQuery(conn data.Connection, query string) (*data.Result, error) {
	result := &data.Result{
		Connection: conn.Name,
		Query:      query,
	}
	db, e := sql.Open(conn.Type, conn.Dsn)
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
		for i, v := range vs {
			s := ""
			if v != nil {
				s = stringify(v)
			}
			if result.Columns[i] == "insert_id" {
				s = ParseUnix(s).String()
			}
			row.Values = append(row.Values, s)
		}
		result.Rows = append(result.Rows, row)
	}

	return result, nil
}

func stringify(v interface{}) string {
	if ints, ok := v.([]uint8); ok {
		return stringify(ints)
	}
	return fmt.Sprint(v)
}
