package outputport

import "time"

func NewUserCreateData(userID string, created time.Time) *UserCreateData {
	return &UserCreateData{
		UserID:  userID,
		Created: created,
	}
}

type UserCreateData struct {
	UserID  string
	Created time.Time
}
