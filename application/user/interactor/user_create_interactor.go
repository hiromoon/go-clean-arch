package interactor

import (
	"github.com/hiromoon/go-clean-arch/application/user/port"
	"github.com/hiromoon/go-clean-arch/domain/model/user"
)

type UserCreateInteractor struct {
	repo user.Repository
}

func NewUserCreateInteractor(repo user.Repository) port.UserCreateInputPort {
	return &UserCreateInteractor{
		repo: repo,
	}
}

func (i *UserCreateInteractor) Handle(input *port.UserCreateInputData) (*port.UserCreateOutputData, error) {
	user := user.NewUser(input.User.ID, input.User.Name, input.User.Password)
	if err := i.repo.Create(user); err != nil {
		return nil, err
	}

	return &port.UserCreateOutputData{}, nil
}
