package main

import (
	"log"
	"net/http"

	"main.go/internal/handlers"
	"main.go/internal/middleware"
	"main.go/internal/router"
	"main.go/internal/storage"
)

func main() {
	db := storage.PostgreSqlInit()
	defer db.DB.Close()
	h := handlers.NewHandler(db)
	mux := http.NewServeMux()
	mux.Handle("/users", router.MethodRouter{
		"GET":  h.GetUsersHandler,
		"POST": h.CreateUserHandler,
	})
	mux.Handle("/users/", router.MethodRouter{
		"GET":    h.GetUserHandler,
		"PUT":    h.UpdateUserHandler,
		"DELETE": h.DeleteUserHandler,
	})
	log.Fatal(http.ListenAndServe(":8080", middleware.JsonContentTypeMiddleware(mux)))
}
