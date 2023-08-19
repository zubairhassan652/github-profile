// Settings package is used to initialized app with default settings
package settings

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-chi/chi"
	"github.com/zubairhassan652/go-gorilla-mux/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	// Its a app level router
	ChiHandler *chi.Mux
}

// sync.Once is used to implement singleton pattern.
var initializer sync.Once

// app instance.
var app *App

// app level db.
var DB *gorm.DB

var Mongo *mongo.Client

func setup() {
	app = new(App)
	app.ChiHandler = chi.NewRouter()
	app.ChiHandler.Use(DBMiddleware)

	app.registerChiRoutes(appListChi())
	// Mongo = app.initMongo()
	DB = app.initDB()
}

func InitApp() *App {
	initializer.Do(setup)

	return app
}

func GetDB() *gorm.DB {
	return DB
}

func GetMongo() *mongo.Client {
	return Mongo
}

func appListChi() map[string]*chi.Mux {
	return map[string]*chi.Mux{
		"users": users.ExposeRoutes(),
	}
}

func (app *App) initDB() *gorm.DB {
	// Open a database connection
	db, err := gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (app *App) initMongo() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Perform operations using client and collection
	// ...
	return client
}

func (app *App) registerChiRoutes(installedApps map[string]*chi.Mux) {
	for _, r := range installedApps {
		app.ChiHandler.Mount("/", r)
	}
}

// project path
// func findProjectRoot() (string, error) {
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		return "", err
// 	}

// 	for {
// 		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
// 			return cwd, nil
// 		}

// 		parentDir := filepath.Dir(cwd)
// 		if parentDir == cwd {
// 			break
// 		}
// 		cwd = parentDir
// 	}

// 	return "", fmt.Errorf("project root not found")
// }
