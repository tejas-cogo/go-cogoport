package users

import (
	u "github.com/tejas-cogo/go-cogoport/helpers"
	"github.com/tejas-cogo/go-cogoport/models"
)

type UserService struct {
	GoUser models.GoUser
}

func UserList() map[string]interface{} {
	userData := models.GetAllUsers
	response := u.Message(0, "This is from version 1 api")
	response["data"] = userData
	return response
}
