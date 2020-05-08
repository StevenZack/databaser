package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/StevenZack/databaser/util"
	"github.com/StevenZack/databaser/views"

	"github.com/StevenZack/databaser/data"

	"github.com/StevenZack/databaser/vars"
)

func Query_POST(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	query := r.FormValue("query")
	c, e := vars.GetConnectionByName(name)
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	switch c.Type {
	case data.TypeClickhouse:
		result, e := util.QueryClickhouse(c, query)
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		t, e := template.New("query").Parse(string(views.Bytes_QueryHtml))
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		e = t.Execute(w, result)
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
	case data.TypeMySQL:
		result, e := util.QueryMysql(c, query)
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		t, e := template.New("query").Parse(string(views.Bytes_QueryHtml))
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		e = t.Execute(w, result)
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "connection type '"+c.Type+"' not support yet", http.StatusBadRequest)
		return
	}
}
