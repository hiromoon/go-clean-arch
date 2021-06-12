package interactor

import (
	"github.com/hiromoon/go-clean-arch/application/user/port"
	"github.com/hiromoon/go-clean-arch/domain/model/user"
)

type UserListInteractor struct {
	repo user.Repository
}

func NewUserListInteractor(repo user.Repository) port.UserListInputPort {
	return &UserListInteractor{repo: repo}
}

func (i *UserListInteractor) Handle(_input *port.UserListInputData) (*port.UserListOutputData, error) {
	users, err := i.repo.GetAll()
	if err != nil {
		return nil, err
	}

	userData := make([]*port.UserData, 0)
	for _, u := range users {
		userData = append(userData, &port.UserData{
			ID:      u.ID,
			Name:     u.Name,
			Password: u.Password,
		})
	}
	return &port.UserListOutputData{
		Users: userData,
	}, nil
}
