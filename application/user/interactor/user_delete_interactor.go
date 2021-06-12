package interactor

import (
	"github.com/hiromoon/go-clean-arch/application/user/port"
	"github.com/hiromoon/go-clean-arch/domain/model/user"
)

type UserDeleteInteractor struct {
	repo user.Repository
}

func NewUserDeleteInteractor(repo user.Repository) port.UserDeleteInputPort {
	return &UserDeleteInteractor{repo: repo}
}

func (i *UserDeleteInteractor) Handle(input *port.UserDeleteInputData) (*port.UserDeleteOutputData, error) {
	user, err := i.repo.Get(input.UserID)
	if err != nil {
		return nil, err
	}

	if err := i.repo.Delete(user.ID); err != nil {
		return nil, err
	}
	return &port.UserDeleteOutputData{}, nil
}
