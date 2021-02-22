package inputport

func NewUserCreateData(userName string) *UserCreateData {
	return &UserCreateData{
		UserName: userName,
	}
}

type UserCreateData struct {
	UserName string
}
