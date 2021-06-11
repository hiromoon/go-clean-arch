package repository

import (
	"github.com/hiromoon/go-api-reference/domain/model/user"
	"github.com/hiromoon/go-api-reference/infra"
)

type UserRepository struct {
	DB *infra.Database
}

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

func NewUserRepository(db *infra.Database) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) GetAll() ([]*user.User, error) {
	rows := []User{}
	if err := repo.DB.DB.Select(&rows, "SELECT * FROM users"); err != nil {
		return nil, err
	}

	users := []*user.User{}
	for _, row := range rows {
		users = append(users, user.NewUser(row.ID, row.Name, row.Password))
	}

	return users, nil
}

func (repo *UserRepository) Create(user *user.User) error {
	_, err := repo.DB.DB.NamedExec("INSERT INTO users (id, name, password) VALUES (:id, :name, :password)", user)
	return err
}

func (repo *UserRepository) Get(id string) (*user.User, error) {
	row := User{}
	if err := repo.DB.DB.Get(&row, "SELECT * FROM users WHERE id=?", id); err != nil {
		return nil, err
	}

	return user.NewUser(row.ID, row.Name, row.Password), nil
}

func (repo *UserRepository) Update(user *user.User) error {
	_, err := repo.DB.DB.NamedExec("UPDATE users SET name=:name, password=:password WHERE id=:id", user)
	return err
}

func (repo *UserRepository) Delete(id string) error {
	_, err := repo.DB.DB.NamedExec("DELETE FROM users WHERE id=:id", map[string]interface{}{"id": id})
	return err
}
