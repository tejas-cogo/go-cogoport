package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/group_members"
)

// func ListGroupMember(c *gin.Context) {
// 	c.JSON(200, service.ListGroupMember())
// }

func CreateGroupMember(c *gin.Context) {
	var group_member models.GroupMember
	c.BindJSON(&group_member)
	c.JSON(200, service.CreateGroupMember(group_member))
}

// func DeleteGroupMember(c *gin.Context) {
// 	id := c.Request.URL.Query().Get("ID")
// 	c.JSON(200, service.DeleteGroupMember(id))
// }

func UpdateGroupMember(c *gin.Context) {
	var body models.GroupMember
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateGroupMember(id, body))
}
