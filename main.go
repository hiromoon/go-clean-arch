package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hiromoon/go-api-reference/controllers"
	"github.com/hiromoon/go-api-reference/infra"
	"github.com/hiromoon/go-api-reference/middlewares"
	"github.com/hiromoon/go-api-reference/repositories"
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
	basicAuthMiddleware := middlewares.NewBasicAuthenticationMiddleware(
		redis,
		repositories.NewUserRepository(db),
	)
	r.Use(basicAuthMiddleware.Middleware)

	usersController := controllers.NewUsersController(
		repositories.NewUserRepository(db),
	)
	r.HandleFunc("/api/v1/users", usersController.UsersCreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/users", usersController.UsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id}", usersController.UserHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id}", usersController.UserUpdateHandler).Methods(http.MethodPatch)
	r.HandleFunc("/api/v1/users/{id}", usersController.UserDeleteHandler).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8000", r))
}
