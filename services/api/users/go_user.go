package users

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type UserService struct {
	GoUser models.GoUser
}

func UserList() []models.GoUser {
	db := config.GetDB()

	var users []models.GoUser
	result := map[string]interface{}{}
	db.Find(&users).Take(&result)

	for _, row := range users {
		fmt.Println("values: ", row.ID, row.Name, "\n")
	}

	return users
}
