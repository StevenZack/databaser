package main

import (
	"fmt"
	"log"

	"github.com/StevenZack/databaser/handler"
	"github.com/StevenZack/databaser/views"
	"github.com/StevenZack/openurl"

	"github.com/StevenZack/mux"
	"github.com/StevenZack/tools/strToolkit"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	port := strToolkit.RandomPort()
	server := mux.NewServer(":" + port)

	server.GET("/", handler.Index)

	server.GET("/connection", handler.Connection_GET)
	server.POST("/query", handler.Query_POST)

	server.HandleHtml("/new", views.Bytes_NewHtml)
	server.POST("/new", handler.New_POST)

	server.GET("/edit", handler.Edit_GET)
	server.POST("/edit", handler.Edit_POST)

	server.GET("/delete", handler.Delete)

	fmt.Println("listen on http://127.0.0.1:" + port)
	openurl.OpenApp("http://127.0.0.1:" + port)
	e := server.ListenAndServe()
	if e != nil {
		fmt.Println("listen error :", e)
		return
	}
}
