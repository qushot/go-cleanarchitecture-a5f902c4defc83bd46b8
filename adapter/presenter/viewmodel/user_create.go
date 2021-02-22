package viewmodel

func NewUserCreate(userID string, createDate string) *userCreate {
	return &userCreate{
		UserID:     userID,
		CreateDate: createDate,
	}
}

type userCreate struct {
	UserID     string
	CreateDate string
}
