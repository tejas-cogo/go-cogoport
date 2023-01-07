package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/roles"
)

// func ListRole(c *gin.Context) {
// 	c.JSON(200, service.ListRole())
//id := c.Request.URL.Query().Get("ID")
// }

func CreateRole(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)
	c.JSON(200, service.CreateRole(role))
}

func DeleteRole(c *gin.Context) {
	var body models.Role
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteRole(id))
}

func UpdateRole(c *gin.Context) {
	var body models.Role
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateRole(id, body))
}
