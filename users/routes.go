// It contains all the `routes` related to users app
package users

import (
	"github.com/go-chi/chi"
	"github.com/zubairhassan652/go-vue/users/handlers"
)

func ExposeRoutes() *chi.Mux {
	mainRouter := chi.NewRouter()
	mainRouter.HandleFunc("/", handlers.HandleHome)

	// Create a separate router for the /api route
	apiRouter := chi.NewRouter()
	apiRouter.HandleFunc("/users", handlers.HandleUsers)
	apiRouter.HandleFunc("/posts", handlers.HandlePosts)

	// Attach the /api router to the main router
	mainRouter.Mount("/api", apiRouter)

	apiRouter1 := chi.NewRouter()
	apiRouter1.HandleFunc("/def", handlers.HandleUsers1)
	apiRouter1.HandleFunc("/geh", handlers.HandlePosts1)

	// Attach the /abc router to the main router
	mainRouter.Mount("/abc", apiRouter1)

	return mainRouter
}
