package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/roles"
)

func ListRole(c *gin.Context) {
	var filters models.Role

	filters.Name = c.Request.URL.Query().Get("filters[name]")

	ser, db := service.ListRole(filters)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

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
	c.JSON(200, service.UpdateRole( body))
}
