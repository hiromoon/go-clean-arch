package user

type Repository interface {
	GetAll() ([]*User, error)
	Create(user *User) error
	Get(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}
