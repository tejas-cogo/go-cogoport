package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/users"
)

func UserList(c *gin.Context) {
	c.JSON(200, service.UserList())
}

func CreateUser(c *gin.Context) {
	var user models.GoUser
	c.BindJSON(&user)
	c.JSON(200, service.CreateUser(user))
}

func DeleteUser(c *gin.Context) {
	id := c.Request.URL.Query().Get("ID")
	c.JSON(200, service.DeleteUser(id))
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, service.UpdateUser(c))
}
