package interactor

import (
	"errors"
	"time"

	"go-clean-architecture-sample/adapter/presenter"
	"go-clean-architecture-sample/adapter/repository"
	"go-clean-architecture-sample/application/usecase"
	"go-clean-architecture-sample/application/usecase/inputport"
	"go-clean-architecture-sample/application/usecase/outputport"
	"go-clean-architecture-sample/entity"
)

func NewUserCreate(userRepository repository.User, presenter presenter.UserCreate) usecase.UserCreate {
	return &userCreate{
		userRepository: userRepository,
		presenter:      presenter,
	}
}

type userCreate struct {
	userRepository repository.User
	presenter      presenter.UserCreate
}

func (u *userCreate) Handle(data *inputport.UserCreateData) error {
	result, err := u.userRepository.FindByUserName(data.UserName)
	if err != nil {
		return errors.New("DBアクセス自体のエラー")
	}

	if result != nil {
		return errors.New("duplicated")
	}

	user := &entity.User{
		UserName: data.UserName,
	}

	if err := u.userRepository.Save(user); err != nil {
		return err
	}

	output := outputport.NewUserCreateData(user.UserID, time.Now())
	u.presenter.Complete(output)

	return nil
}
