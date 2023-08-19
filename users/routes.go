// It contains all the `routes` related to users app
package users

import (
	"github.com/gorilla/mux"
	"github.com/zubairhassan652/go-gorilla-mux/users/handlers"
)

func ExposeRoutes() *mux.Router {
	mainRouter := mux.NewRouter()

	// Attach routes and handlers to the main router
	mainRouter.HandleFunc("/", handlers.HandleHome)

	// Create a separate router for the /api route
	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/users", handlers.HandleUsers)
	apiRouter.HandleFunc("/posts", handlers.HandlePosts)

	// Attach the /api router to the main router
	mainRouter.PathPrefix("/api").Handler(apiRouter)

	// Gorilla does not support route duplication
	// Create a separate router for the /api route
	apiRouter1 := mux.NewRouter().PathPrefix("/abc").Subrouter()
	apiRouter1.HandleFunc("/def", handlers.HandleUsers1)
	apiRouter1.HandleFunc("/geh", handlers.HandlePosts1)
	// Attach the /api router to the main router
	mainRouter.PathPrefix("/abc").Handler(apiRouter1)

	return mainRouter
}
