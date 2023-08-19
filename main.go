package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zubairhassan652/go-gorilla-mux/settings"
)

func main() {
	app := settings.InitApp()
	fs := http.FileServer(http.Dir("static"))
	app.Handler.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", app.Handler)
	fmt.Println("App listening at 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
