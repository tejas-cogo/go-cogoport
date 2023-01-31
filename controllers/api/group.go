package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/groups"
)

func ListGroup(c *gin.Context) {
	var filters models.FilterGroup
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(400, "Not Found")
	}

	ser, db, err := service.ListGroup(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func ListGroupTag(c *gin.Context) {
	var Tag string
	Tag = c.Request.URL.Query().Get("Tag")
	ser, db, err := service.ListGroupTag(Tag)
	if err != nil {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateGroup(c *gin.Context) {
	var group models.Group
	c.BindJSON(&group)
	ser, err := service.CreateGroup(group)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteGroup(c *gin.Context) {
	var body models.Group
	c.BindJSON(&body)
	id := body.ID
	ser, err := service.DeleteGroup(id)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateGroup(c *gin.Context) {
	var body models.Group
	c.BindJSON(&body)
	ser, err := service.UpdateGroup(body)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
