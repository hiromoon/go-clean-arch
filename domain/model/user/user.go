package user

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

func (u *User) ChangeName(name string) {
	u.Name = name
}

func (u *User) ChangePassword(password string) {
	u.Password = password
}
