package controller

import (
	"encoding/json"
	"github.com/hiromoon/go-api-reference/application/user/port"
	"github.com/hiromoon/go-api-reference/domain/model/user"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UsersController struct {
	repo               user.Repository
	userListInteractor port.UserListInputPort
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

func NewUsersController(
	repo user.Repository,
	userListInteractor port.UserListInputPort,
) *UsersController {
	return &UsersController{
		repo:               repo,
		userListInteractor: userListInteractor,
	}
}

func (c *UsersController) Index(w http.ResponseWriter, r *http.Request) {
	output, err := c.userListInteractor.Handle(&port.UserListInputData{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	users := make([]*User, 0)
	for _, u := range output.Users {
		users = append(users, &User{ID: u.ID, Name: u.Name, Password: u.Password})
	}

	responsePayload := &UsersResponsePayload{Users: users}
	if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UsersController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := c.repo.Get(vars["id"])
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

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var reqPayload UsersCreateRequestPayload
	if err := json.Unmarshal(body, &reqPayload); err != nil {
		log.Fatal(err)
	}

	user := user.NewUser(reqPayload.User.ID, reqPayload.User.Name, reqPayload.User.Password)
	if err := c.repo.Create(user); err != nil {
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

func (c *UsersController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := c.repo.Get(vars["id"])
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

	if err := c.repo.Update(user); err != nil {
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

func (c *UsersController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := c.repo.Delete(vars["id"]); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
}
