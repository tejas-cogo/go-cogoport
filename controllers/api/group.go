package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/groups"
)

func ListGroup(c *gin.Context) {
	var filters models.FilterGroup
	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	ser, db := service.ListGroup(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func ListGroupTag(c *gin.Context) {
	var Tag string
	Tag = c.Request.URL.Query().Get("Tag")
	c.JSON(c.Writer.Status(), service.ListGroupTag(Tag))
}

func CreateGroup(c *gin.Context) {
	var group models.Group
	c.BindJSON(&group)
	ser, err := service.CreateGroup(group)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else if ser != "Successfully Created!" {
		c.JSON(400, ser)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteGroup(c *gin.Context) {
	var body models.Group
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteGroup(id))
}

func UpdateGroup(c *gin.Context) {
	var body models.Group
	c.BindJSON(&body)
	c.JSON(200, service.UpdateGroup(body))
}
