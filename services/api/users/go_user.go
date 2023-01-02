package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

func UpdateUser(c *gin.Context) models.GoUser {
	db := config.GetDB()

	var user models.GoUser

	body := models.GoUser{}
	c.BindJSON(&body)

	id := c.Request.URL.Query().Get("ID")
	db.Where("id = ?", id).First(&user)

	user.Name = body.Name

	db.Save(&user)
	return user
}
