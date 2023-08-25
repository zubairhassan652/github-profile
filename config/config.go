// Config package is used to initialize app with default settings
package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type env struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBSslMode  string `mapstructure:"DB_SSL_MODE"`
	MongoDBUri string `mapstructure:"MONGO_DB_URI"`
}

type WebConfig struct {
	// Its a app level router
	Router *chi.Mux

	// Configure app level dbs
	SqlClient   *gorm.DB
	MongoClient *mongo.Client

	// Environment variables of the app
	Envs env
}

// sync.Once is used to implement singleton pattern.
var initializer sync.Once

// app instance.
var App *WebConfig

func setup() {
	App = new(WebConfig)
	App.loadEnv()
	App.Router = chi.NewRouter()
	App.MongoClient = App.initMongoClient()
	App.SqlClient = App.initSqlClient()
}

func InitApp() *WebConfig {
	initializer.Do(setup)

	return App
}

func (app *WebConfig) initSqlClient() *gorm.DB {
	// Replace with your PostgreSQL database URL
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		app.Envs.DBHost,
		app.Envs.DBPort,
		app.Envs.DBUser,
		app.Envs.DBName,
		app.Envs.DBSslMode,
		app.Envs.DBPassword,
	)
	// dbURL := "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=your-postgres-password"

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to Postgres")

	return db
}

func (app *WebConfig) initMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(app.Envs.MongoDBUri)
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

// func DBMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		db := App.SqlClient
// 		// client := Mongo
// 		ctx := context.WithValue(req.Context(), "db", db)
// 		// ctx := context.WithValue(req.Context(), "client", client)
// 		req = req.WithContext(ctx)
// 		next.ServeHTTP(res, req)
// 	})
// }

// Load env vars from .env file.
func (app *WebConfig) loadEnv() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("env not found")
	}
	// Set envs from viper to os
	for _, key := range viper.AllKeys() {
		os.Setenv(strings.ToUpper(key), viper.GetString(key))
	}
	// fmt.Println(viper.AllKeys())
	// fmt.Println(os.Getenv("DB_HOST"))
	// Unmarshal configurations into struct
	if err := viper.Unmarshal(&app.Envs); err != nil {
		fmt.Printf("Error unmarshaling config: %s\n", err)
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
