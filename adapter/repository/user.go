package repository

import (
	"go-clean-architecture-sample/entity"
)

type User interface {
	FindByUserName(userName string) (*entity.User, error)
	Save(user *entity.User) error
}
