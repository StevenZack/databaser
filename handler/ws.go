package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Ws(w http.ResponseWriter, r *http.Request) {
	c, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	defer c.Close()
	for {
		_, b, e := c.ReadMessage()
		if e != nil {
			log.Println(e)
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(string(b))
	}
}
