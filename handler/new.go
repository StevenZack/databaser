package handler

import (
	"log"
	"net/http"
	"net/url"

	"github.com/StevenZack/databaser/vars"

	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"

	"github.com/StevenZack/databaser/data"
	"github.com/go-sql-driver/mysql"
)

func New_Post(w http.ResponseWriter, r *http.Request) {
	c := data.Connection{
		Name: r.FormValue("name"),
		Type: r.FormValue("type"),
		Dsn:  r.FormValue("dsn"),
	}
	switch c.Type {
	case data.TypeMySQL:
		info, e := mysql.ParseDSN(c.Dsn)
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		c.DB = info.DBName
	case data.TypeClickhouse:
		info, e := url.Parse(c.Dsn)
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		c.DB = info.Query().Get("database")
	case data.TypeMongoDB:
		info, e := connstring.Parse(c.Dsn)
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		c.DB = info.Database
	default:
		http.Error(w, "unsupported connection type '"+c.Type+"'", http.StatusInternalServerError)
		return
	}
	vars.AppendConnection(c)
	http.Redirect(w, r, "/", http.StatusFound)
}
