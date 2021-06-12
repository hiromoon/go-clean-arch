package interactor

import (
	"github.com/hiromoon/go-clean-arch/application/user/port"
	"github.com/hiromoon/go-clean-arch/domain/model/user"
)

type UserFindInteractor struct {
	repo user.Repository
}

func NewUserFindInteractor(repo user.Repository) *UserFindInteractor {
	return &UserFindInteractor{
		repo: repo,
	}
}

func (i *UserFindInteractor) Handle(input *port.UserFindInputData) (*port.UserFindOutputData, error) {
	user, err := i.repo.Get(input.UserID)
	if err != nil {
		return nil, err
	}
	return &port.UserFindOutputData{
		User: &port.UserData{
			ID:       user.ID,
			Name:     user.Name,
			Password: user.Password,
		},
	}, nil
}
