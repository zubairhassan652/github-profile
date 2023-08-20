package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zubairhassan652/go-vue/config"
)

func main() {
	app := config.InitApp()
	fs := http.FileServer(http.Dir("static"))
	app.ChiHandler.Handle("/static/*", http.StripPrefix("/static/", fs))
	http.Handle("/", app.ChiHandler)

	fmt.Println("App listening at 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
