package users

import (
	"github.com/tejas-cogo/go-cogoport/models"
)

type UserService struct {
	GoUser models.GoUser
}

func (us *UserService) UserList() map[string]interface{} {
	userData := us.GoUser.GetAllUsers
	// response := u.Message(0, "This is from version 1 api")
	// response["data"] = userData
	return userData
}
