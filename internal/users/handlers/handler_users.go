package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/zubairhassan652/go-vue/config"
	"github.com/zubairhassan652/go-vue/internal/users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func HandlePostgres(w http.ResponseWriter, r *http.Request) {
	db := config.App.SqlClient
	// db, ok := r.Context().Value("db").(*gorm.DB)
	// if !ok {
	// 	fmt.Println("db not found")
	// }

	fmt.Println("gorm db:", db)

	age := 32
	user := models.User{
		Name: "Zubair",
		Age:  age,
	}
	db.Create(&user)
	db.Model(models.User{}).Where("id < ?", user.ID).Updates(models.User{
		Name: "zubair",
	})

	allUsers := []models.User{}
	db.Find(&allUsers)

	if maxAllowedUsers := 3; len(allUsers) > maxAllowedUsers {
		firstUser := models.User{}
		db.Model(models.User{}).First(&firstUser).Delete(firstUser) //.Statement.SQL
	}

	// response
	response, _ := json.Marshal(user)

	// fmt.Fprintln(w, response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func HandleMongoDB(w http.ResponseWriter, r *http.Request) {
	client := config.App.MongoClient

	// Insert document.
	collection := client.Database("mydb").Collection("mycollection")
	data := bson.D{{Key: "key", Value: "value of key"}}
	_, err = collection.InsertOne(context.Background(), data)

	if err != nil {
		log.Fatal(err)
	}

	// Sort the documents in descending order based on a timestamp field (replace "timestamp" with your field name)
	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}})

	// Querying the document.
	cursor, err := collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Delete record if they are more than three and write first ones to the response writer.
	mongodbRecordsUpdates(w, cursor, collection)
	w.Header().Set("Content-Type", "application/json")
}

func mongodbRecordsUpdates(w http.ResponseWriter, cursor *mongo.Cursor, collection *mongo.Collection) {
	numberOfRecords, limit := 0, 3

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)

		if numberOfRecords < limit {
			// response
			response, err := json.Marshal(result)

			if err != nil {
				log.Println(err)
			}

			w.Write(response)
		} else {
			// Deleting the document.
			_, err = collection.DeleteOne(context.Background(), result)

			if err != nil {
				log.Fatal(err)
			}
		}
		numberOfRecords++
	}
}
