package handler

import (
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/StevenZack/databaser/data"
	"github.com/StevenZack/databaser/vars"
	"github.com/StevenZack/databaser/views"
	"github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func Edit_GET(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	c, e := vars.GetConnectionByName(name)
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	t, e := template.New("edit").Parse(string(views.Bytes_EditHtml))
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	e = t.Execute(w, c)
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}

func Edit_POST(w http.ResponseWriter, r *http.Request) {
	originName := r.FormValue("origin_name")
	_, e := vars.GetConnectionByName(originName)
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

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

	vs := vars.GetConnections()
	for i, v := range vs {
		if v.Name == originName {
			cs := append(vs[:i], append([]data.Connection{c}, vs[i+1:]...)...)
			vars.SetConnections(cs)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
	http.Error(w, "connection '"+originName+"' not found", http.StatusNotFound)
}
