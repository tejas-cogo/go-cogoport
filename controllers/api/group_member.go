package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/group_members"
)

func ListGroupMember(c *gin.Context) {
	var group_member models.GroupMember
	TicketID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[group_id]"))
	group_member.GroupID = uint(TicketID)
	ser, db := service.ListGroupMember(group_member)

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
	c.JSON(200, service.CreateGroupMember(group_member))
}

func DeleteGroupMember(c *gin.Context) {
	var body models.GroupMember
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteGroupMember(id))
}

func UpdateGroupMember(c *gin.Context) {
	var body models.GroupMember
	c.BindJSON(&body)
	c.JSON(200, service.UpdateGroupMember(body))
}
