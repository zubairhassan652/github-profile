// It contains all the `routes` related to users app
package users

import (
	"github.com/go-chi/chi"
	"github.com/zubairhassan652/go-vue/internal/users/handlers"
)

func Routes() *chi.Mux {
	mainRouter := chi.NewRouter()
	mainRouter.HandleFunc("/", handlers.HandleHome)

	// Create a separate router for the /api route
	apiRouter := chi.NewRouter()
	apiRouter.HandleFunc("/users", handlers.HandleUsers)
	apiRouter.HandleFunc("/posts", handlers.HandlePosts)

	// Attach the /api router to the main router
	mainRouter.Mount("/api", apiRouter)

	apiRouter1 := chi.NewRouter()
	apiRouter1.HandleFunc("/postgres", handlers.HandlePostgres)
	apiRouter1.HandleFunc("/mongodb", handlers.HandleMongoDB)

	// Attach the /abc router to the main router
	mainRouter.Mount("/db", apiRouter1)

	return mainRouter
}
