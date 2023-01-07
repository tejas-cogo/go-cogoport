package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_task_assignees"
)

func ListTicketTaskAssignee(c *gin.Context) {
	c.JSON(200, service.ListTicketTaskAssignee())
}

func CreateTicketTaskAssignee(c *gin.Context) {
	var ticket_task_assignee models.TicketTaskAssignee
	c.BindJSON(&ticket_task_assignee)
	c.JSON(200, service.CreateTicketTaskAssignee(ticket_task_assignee))
}

// func DeleteTicketTaskAssignee(c *gin.Context) {
// 	id := c.Request.URL.Query().Get("ID")
// 	c.JSON(200, service.DeleteTicketTaskAssignee(id))
// }

func UpdateTicketTaskAssignee(c *gin.Context) {
	var body models.TicketTaskAssignee
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketTaskAssignee(id, body))
}
