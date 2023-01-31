package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/roles"
)

func ListRole(c *gin.Context) {
	var filters models.Role

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(400, "Not Found")
	}

	ser, db, err := service.ListRole(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateRole(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)
	ser, err := service.CreateRole(role)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteRole(c *gin.Context) {
	var body models.Role
	c.BindJSON(&body)
	id := body.ID
	ser, err := service.DeleteRole(id)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateRole(c *gin.Context) {
	var body models.Role
	c.BindJSON(&body)
	ser, err := service.UpdateRole(body)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
