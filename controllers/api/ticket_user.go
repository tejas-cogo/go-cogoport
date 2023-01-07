package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
)

func ListTicketUser(c *gin.Context) {
	c.JSON(200, service.ListTicketUser())
}

func CreateTicketUser(c *gin.Context) {
	var ticket_user models.TicketUser
	c.BindJSON(&ticket_user)
	c.JSON(200, service.CreateTicketUser(ticket_user))
}

// func DeleteTicketUser(c *gin.Context) {
// 	id := c.Request.URL.Query().Get("ID")
// 	c.JSON(200, service.DeleteTicketUser(id))
// }

func UpdateTicketUser(c *gin.Context) {
	var body models.TicketUser
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketUser(id, body))
}
