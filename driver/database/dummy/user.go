package dummy

import (
	"fmt"

	"go-clean-architecture-sample/adapter/repository"
	"go-clean-architecture-sample/entity"
)

var _ repository.User = &user{}

func NewUser() repository.User {
	return &user{}
}

type user struct{}

func (u *user) FindByUserName(userName string) (*entity.User, error) {
	fmt.Printf("[===== userName(%s)を使って検索する処理 =====]\n", userName)

	// エラーも出なかったしユーザもいなかった想定
	return nil, nil
}

func (u *user) Save(user *entity.User) error {
	fmt.Printf("[===== user(%v)を永続化する処理 =====]\n", user)

	user.UserID = "12345"

	// 保存が成功した想定
	return nil
}
