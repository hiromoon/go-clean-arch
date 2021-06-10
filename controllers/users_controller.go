package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hiromoon/go-api-reference/models"
	"github.com/hiromoon/go-api-reference/repositories"
)

type UsersController struct {
	repository *repositories.UserRepository
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UsersCreateRequestPayload struct {
	User *User `json:"user"`
}

type UsersCreateResponsePayload struct {
	User *User `json:"user"`
}

type UsersResponsePayload struct {
	Users []*User `json:"users"`
}

type UserResponsePayload struct {
	User *User `json:"user"`
}

type UserUpdateRequestPayload struct {
	User *User `json:"user"`
}

type UserUpdateResponsePayload struct {
	User *User `json:"user"`
}

func NewUsersController(repository *repositories.UserRepository) *UsersController {
	return &UsersController{
		repository: repository,
	}
}

func (c *UsersController) UsersCreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var reqPayload UsersCreateRequestPayload
	if err := json.Unmarshal(body, &reqPayload); err != nil {
		log.Fatal(err)
	}

	user := models.NewUser(reqPayload.User.ID, reqPayload.User.Name, reqPayload.User.Password)
	if err := c.repository.Create(user); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	responsePayload := &UsersCreateResponsePayload{User: &User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}}
	json.NewEncoder(w).Encode(responsePayload)
}

func (c *UsersController) UsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := c.repository.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	us := []*User{}
	for _, u := range users {
		us = append(us, &User{ID: u.ID, Name: u.Name, Password: u.Password})
	}
	responsePayload := &UsersResponsePayload{Users: us}
	json.NewEncoder(w).Encode(responsePayload)
}

func (c *UsersController) UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := c.repository.Get(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responsePayload := &UserResponsePayload{User: &User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}}
	json.NewEncoder(w).Encode(responsePayload)
}

func (c *UsersController) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := c.repository.Get(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var reqPayload UserUpdateRequestPayload
	if err := json.Unmarshal(body, &reqPayload); err != nil {
		log.Fatal(err)
	}

	user.ID = reqPayload.User.ID
	user.Name = reqPayload.User.Name
	user.Password = reqPayload.User.Password

	if err := c.repository.Update(user); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responsePayload := &UserUpdateResponsePayload{User: &User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}}
	json.NewEncoder(w).Encode(responsePayload)
}

func (c *UsersController) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := c.repository.Delete(vars["id"]); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
}
