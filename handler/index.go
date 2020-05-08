package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/StevenZack/databaser/vars"
	"github.com/StevenZack/databaser/views"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, e := template.New("index").Parse(string(views.Bytes_IndexHtml))
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	e = t.Execute(w, vars.GetConnections())
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}
