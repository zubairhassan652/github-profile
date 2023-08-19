package settings

import (
	"context"
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func DBMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		db := DB
		// client := Mongo
		ctx := context.WithValue(req.Context(), "db", db)
		// ctx := context.WithValue(req.Context(), "client", client)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}
