package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/group_members"
)

type Test struct {
	ID uint `query:"id"`
}

func ListGroupMember(c *gin.Context) {
	var filters models.FilterGroupMember
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	if err != nil {
		c.JSON(c.Writer.Status(), "Not Found")
	}

	ser, db := service.ListGroupMember(filters)

	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateGroupMember(c *gin.Context) {
	var group_member models.CreateGroupMember
	err := c.Bind(&group_member)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.CreateGroupMember(group_member)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteGroupMember(c *gin.Context) {
	var body models.GroupMember
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	id := body.ID

	ser, err := service.DeleteGroupMember(id)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateGroupMember(c *gin.Context) {
	var body models.GroupMember
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.UpdateGroupMember(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
