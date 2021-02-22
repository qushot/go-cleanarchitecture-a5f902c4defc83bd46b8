package presenter

import (
	"fmt"

	"go-clean-architecture-sample/adapter/presenter/viewmodel"
	"go-clean-architecture-sample/application/usecase/outputport"
)

var _ UserCreate = &userCreate{}

type UserCreate interface {
	Complete(data *outputport.UserCreateData)
}

func NewUserCreate() UserCreate {
	return &userCreate{}
}

type userCreate struct{}

func (u *userCreate) Complete(data *outputport.UserCreateData) {
	createdDateStr := data.Created.Format("2006/01/02")
	model := viewmodel.NewUserCreate(data.UserID, createdDateStr)
	fmt.Printf("id: %s, creted: %s\n", model.UserID, model.CreateDate)
}
