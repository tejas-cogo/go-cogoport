package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/group_members"
)

func ListGroupMember(c *gin.Context) {
	var group_member models.GroupMember
	c.BindJSON(&group_member)
	ser, db := service.ListGroupMember(group_member)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func CreateGroupMember(c *gin.Context) {
	var group_member models.GroupMember
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
	var body models.Filter
	c.BindJSON(&body)
	c.JSON(200, service.UpdateGroupMember(body))
}
