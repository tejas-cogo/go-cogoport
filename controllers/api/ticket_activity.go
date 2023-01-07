package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

// func ListTicketActivity(c *gin.Context) {
// 	c.JSON(200, service.TicketActivityList())
// }

func CreateTicketActivity(c *gin.Context) {
	var ticket_activity models.TicketActivity
	c.BindJSON(&ticket_activity)
	c.JSON(200, service.CreateTicketActivity(ticket_activity))
}

// func DeleteTicketActivity(c *gin.Context) {
// 	id := c.Request.URL.Query().Get("ID")
// 	c.JSON(200, service.DeleteTicketActivity(id))
// }

func UpdateTicketActivity(c *gin.Context) {
	var body models.TicketActivity
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketActivity(id, body))
}
