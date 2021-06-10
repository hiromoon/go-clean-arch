package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hiromoon/go-api-reference/controller"
	"github.com/hiromoon/go-api-reference/infra"
	"github.com/hiromoon/go-api-reference/repository"
)

func main() {
	db, err := infra.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	redis := infra.ConnectRedis()
	defer redis.Close()

	r := mux.NewRouter()
	// basicAuthMiddleware := middleware.NewBasicAuthenticationMiddleware(
	// 	redis,
	// 	repository.NewUserRepository(db),
	// )
	// r.Use(basicAuthMiddleware.Middleware)

	usersController := controller.NewUsersController(
		repository.NewUserRepository(db),
	)
	r.HandleFunc("/api/v1/users", usersController.UsersCreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/users", usersController.UsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id}", usersController.UserHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id}", usersController.UserUpdateHandler).Methods(http.MethodPatch)
	r.HandleFunc("/api/v1/users/{id}", usersController.UserDeleteHandler).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8000", r))
}
