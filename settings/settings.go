// Settings package is used to initialized app with default settings
package settings

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/mux"
	"github.com/zubairhassan652/go-gorilla-mux/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	// Its a app level router
	Handler *mux.Router

	// List of user defined apps
	// installedApps map[string]*mux.Router

}

var initializer sync.Once

var app *App

// app level db.
var DB *gorm.DB

var Mongo *mongo.Client

func setup() {
	app = new(App)
	app.Handler = mux.NewRouter()
	app.Handler.Use(DBMiddleware)
	app.registerRoutes(appList())
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

func appList() map[string]*mux.Router {
	return map[string]*mux.Router{
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

// func (app *App) getRegisteredApps() map[string]*mux.Router {
// 	return app.installedApps
// }

// func (app *App) setRegisteredApps(installedApps map[string]*mux.Router) {
// 	app.installedApps = installedApps
// }

// func (app *App) registerRoute(path string, handler http.HandlerFunc) {
// 	app.Handler.HandleFunc(path, handler)
// }

func (app *App) registerRoutes(installedApps map[string]*mux.Router) {
	for _, r := range installedApps {
		app.Handler.NewRoute().Handler(r)
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

// findAppFolderAndRoutes find routes.go file in the app
// func findAppFolderAndRoutes(root string) (string, error) {
// 	var foundPath string

// 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if info.IsDir() && strings.ToLower(info.Name()) == "users" {
// 			routesFilePath := filepath.Join(path, "routes.go")
// 			if _, err := os.Stat(routesFilePath); err == nil {
// 				foundPath = routesFilePath
// 				return filepath.SkipDir // Skip searching further within this directory
// 			}
// 		}

// 		return nil
// 	})

// 	return foundPath, err
// }

// func getRegisteredRoutesFromApp(name string) {
// 	root, err := findProjectRoot()
// 	CheckError(err)
// 	file, err := findAppFolderAndRoutes(root)
// 	CheckError(err)

// 	appRouter := ""
// }
