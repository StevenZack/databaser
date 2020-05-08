package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/StevenZack/databaser/vars"
	"github.com/StevenZack/databaser/views"
)

func Connection_GET(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	c, e := vars.GetConnectionByName(name)
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	t, e := template.New("connection").Parse(string(views.Bytes_ConnectionHtml))
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
