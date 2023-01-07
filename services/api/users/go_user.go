package users

import (
	"fmt"

	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
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

func CreateUser(user models.GoUser) models.GoUser {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&user)
	return user
}

func DeleteUser(id string) models.GoUser {
	db := config.GetDB()
	var user models.GoUser
	db.Where("id = ?", id).Delete(&user)
	return user
}

func UpdateUser(id uint, body models.GoUser) models.GoUser {
	db := config.GetDB()
	var user models.GoUser
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&user)

	user.Name = body.Name

	db.Save(&user)
	return user
}
