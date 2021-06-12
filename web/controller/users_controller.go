package controller

import (
	"encoding/json"
	"github.com/hiromoon/go-api-reference/application/user/port"
	"github.com/hiromoon/go-api-reference/domain/model/user"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type UsersController struct {
	repo                 user.Repository
	userListInteractor   port.UserListInputPort
	userFindInteractor   port.UserFindInputPort
	userCreateInteractor port.UserCreateInputPort
	userUpdateInteractor port.UserUpdateInputPort
	userDeleteInteractor port.UserDeleteInputPort
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
	userFindInteractor port.UserFindInputPort,
	userCreateInteractor port.UserCreateInputPort,
	userUpdateInteractor port.UserUpdateInputPort,
	userDeleteInteractor port.UserDeleteInputPort,
) *UsersController {
	return &UsersController{
		repo:               repo,
		userListInteractor: userListInteractor,
		userFindInteractor: userFindInteractor,
		userCreateInteractor: userCreateInteractor,
		userUpdateInteractor: userUpdateInteractor,
		userDeleteInteractor: userDeleteInteractor,
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

	output, err := c.userFindInteractor.Handle(&port.UserFindInputData{UserID: vars["id"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responsePayload := &UserResponsePayload{User: &User{
		ID:       output.User.ID,
		Name:     output.User.Name,
		Password: output.User.Password,
	}}
	if err := json.NewEncoder(w).Encode(responsePayload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var reqPayload UsersCreateRequestPayload
	if err := json.Unmarshal(body, &reqPayload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = c.userCreateInteractor.Handle(&port.UserCreateInputData{
		User: &port.UserData{
			ID:       reqPayload.User.ID,
			Name:     reqPayload.User.Name,
			Password: reqPayload.User.Password,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (c *UsersController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var reqPayload UserUpdateRequestPayload
	if err := json.Unmarshal(body, &reqPayload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = c.userUpdateInteractor.Handle(&port.UserUpdateInputData{User: &port.UserData{
		ID:       vars["id"],
		Name:     reqPayload.User.Name,
		Password: reqPayload.User.Password,
	}})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (c *UsersController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, err := c.userDeleteInteractor.Handle(&port.UserDeleteInputData{UserID: vars["id"]}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
}
