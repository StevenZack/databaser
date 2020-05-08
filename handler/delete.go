package handler

import (
	"net/http"

	"github.com/StevenZack/databaser/vars"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	cs := vars.GetConnections()
	for i, c := range cs {
		if c.Name == name {
			vars.SetConnections(append(cs[:i], cs[i+1:]...))
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
}
