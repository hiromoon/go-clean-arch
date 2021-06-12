package main

import (
	"github.com/hiromoon/go-api-reference/application/user/interactor"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hiromoon/go-api-reference/infra"
	"github.com/hiromoon/go-api-reference/web/controller"
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

	userRepository := infra.NewUserRepository(db)
	usersController := controller.NewUsersController(
		userRepository,
		interactor.NewUserListInteractor(userRepository),
		interactor.NewUserFindInteractor(userRepository),
	)
	r.HandleFunc("/api/v1/users", usersController.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/users", usersController.Index).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id}", usersController.Show).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id}", usersController.Update).Methods(http.MethodPatch)
	r.HandleFunc("/api/v1/users/{id}", usersController.Delete).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8000", r))
}
