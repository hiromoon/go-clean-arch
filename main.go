package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/dig"

	"github.com/hiromoon/go-clean-arch/application/user/interactor"
	"github.com/hiromoon/go-clean-arch/infra"
	"github.com/hiromoon/go-clean-arch/web/controller"
)

func main() {
	db, err := infra.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	redis := infra.ConnectRedis()
	defer redis.Close()

	c := dig.New()
	if err := c.Provide(func() *infra.Database { return db }); err != nil {
		panic(err)
	}
	if err := c.Provide(infra.NewUserRepository); err != nil {
		panic(err)
	}
	if err := c.Provide(interactor.NewUserListInteractor); err != nil {
		panic(err)
	}
	if err := c.Provide(interactor.NewUserFindInteractor); err != nil {
		panic(err)
	}
	if err := c.Provide(interactor.NewUserCreateInteractor); err != nil {
		panic(err)
	}
	if err := c.Provide(interactor.NewUserUpdateInteractor); err != nil {
		panic(err)
	}
	if err := c.Provide(interactor.NewUserDeleteInteractor); err != nil {
		panic(err)
	}
	if err := c.Provide(controller.NewUsersController); err != nil {
		panic(err)
	}

	var uc *controller.UsersController
	if err := c.Invoke(func(usersController *controller.UsersController) {
		uc = usersController
	}); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	// basicAuthMiddleware := middleware.NewBasicAuthenticationMiddleware(
	// 	redis,
	// 	repository.NewUserRepository(db),
	// )
	// r.Use(basicAuthMiddleware.Middleware)

	r.HandleFunc("/api/v1/users", uc.Create).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/users", uc.Index).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id}", uc.Show).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id}", uc.Update).Methods(http.MethodPatch)
	r.HandleFunc("/api/v1/users/{id}", uc.Delete).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8000", r))
}
