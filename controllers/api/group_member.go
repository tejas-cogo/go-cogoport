package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/group_members"
)

type Test struct {
	ID uint `query:"id"`
}

func ListGroupMember(c *gin.Context) {
	var filters models.FilterGroupMember
	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	}

	ser, db := service.ListGroupMember(filters)

	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateGroupMember(c *gin.Context) {
	var group_member models.CreateGroupMember
	c.BindJSON(&group_member)
	ser, err := service.CreateGroupMember(group_member)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else if ser != "Successfully Created!" {
		c.JSON(c.Writer.Status(), ser)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteGroupMember(c *gin.Context) {
	var body models.GroupMember
	c.BindJSON(&body)
	id := body.ID
	err, ser := service.DeleteGroupMember(id)
	if err != "Successfully Deleted!" {
		c.JSON(c.Writer.Status(), ser)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateGroupMember(c *gin.Context) {
	var body models.GroupMember
	c.BindJSON(&body)
	c.JSON(c.Writer.Status(), service.UpdateGroupMember(body))
}
