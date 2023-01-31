package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/roles"
)

func ListRole(c *gin.Context) {
	var filters models.Role

	// filters.Name = c.Request.URL.Query().Get("filters[name]")

	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	}

	ser, db := service.ListRole(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
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
	} else if ser != "Successfully Created!" {
		c.JSON(c.Writer.Status(), ser)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteRole(c *gin.Context) {
	var body models.Role
	c.BindJSON(&body)
	id := body.ID
	c.JSON(c.Writer.Status(), service.DeleteRole(id))
}

func UpdateRole(c *gin.Context) {
	var body models.Role
	c.BindJSON(&body)
	c.JSON(c.Writer.Status(), service.UpdateRole(body))
}
