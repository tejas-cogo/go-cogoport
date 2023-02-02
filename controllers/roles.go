package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/roles"
)

func ListRole(c *gin.Context) {
	var filters models.Role

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
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
	err := c.Bind(&role)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.CreateRole(role)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteRole(c *gin.Context) {
	var body models.Role
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	id := body.ID
	ser, err := service.DeleteRole(id)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateRole(c *gin.Context) {
	var body models.Role
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.UpdateRole(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
