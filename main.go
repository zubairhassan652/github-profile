package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zubairhassan652/go-vue/config"
	"github.com/zubairhassan652/go-vue/internal/users"
)

func main() {
	app := config.InitApp()
	fs := http.FileServer(http.Dir("static"))
	app.Router.Handle("/static/*", http.StripPrefix("/static/", fs))
	RegisterRouters(app)
	http.Handle("/", app.Router)

	fmt.Println("App listening at 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func RegisterRouters(app *config.WebConfig) {
	// Register all the apps here
	app.Router.Mount("/", users.Routes())
}
