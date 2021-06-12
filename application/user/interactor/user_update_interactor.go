package interactor

import (
	"github.com/hiromoon/go-clean-arch/application/user/port"
	"github.com/hiromoon/go-clean-arch/domain/model/user"
)

type UserUpdateInteractor struct {
	repo user.Repository
}

func NewUserUpdateInteractor(repo user.Repository) port.UserUpdateInputPort {
	return &UserUpdateInteractor{
		repo: repo,
	}
}

func (i *UserUpdateInteractor) Handle(input *port.UserUpdateInputData) (*port.UserUpdateOutputData, error) {
	user, err := i.repo.Get(input.User.ID)
	if err != nil {
		return nil, err
	}

	user.ChangeName(input.User.Name)
	user.ChangePassword(input.User.Password)

	if err := i.repo.Update(user); err != nil {
		return nil, err
	}
	return &port.UserUpdateOutputData{}, nil
}
