package models

type User struct {
	ID       string
	Name     string
	Password string
}

func NewUser(id string, name string, password string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Password: password,
	}
}
