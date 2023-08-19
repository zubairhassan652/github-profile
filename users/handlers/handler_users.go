package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"gorm.io/gorm"
)

var (
	Templates *template.Template
	err       error
)

func initTemplates() {
	templateFiles, err := filepath.Glob("./static/templates/*.html")
	if err != nil {
		log.Fatal("Error finding template files:", err)
	}

	Templates, err = template.ParseFiles(templateFiles...)
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("db").(*gorm.DB)
	if !ok {
		fmt.Println("db not found")
	}

	fmt.Println("gorm db:", db)

	initTemplates()

	data := struct {
		Title   string
		Content string
	}{
		Title:   "Page Title",
		Content: "This is the content of the page.",
	}

	w.Header().Set("Content-Type", "text/html")
	err = Templates.ExecuteTemplate(w, "index.html", data)

	if err != nil {
		http.Error(w, "Error rendering HTML template", http.StatusInternalServerError)
	}
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of users")
}

func HandlePosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of posts")
}

func HandleUsers1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of users1")
}

func HandlePosts1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of posts1")
}
