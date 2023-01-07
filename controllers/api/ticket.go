package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

// func ListTicket(c *gin.Context) {
// 	c.JSON(200, service.ListTicket())
// }

func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	c.BindJSON(&ticket)
	c.JSON(200, service.CreateTicket(ticket))
}

func DeleteTicket(c *gin.Context) {
	var body models.Ticket
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicket(id))
}

func UpdateTicket(c *gin.Context) {
	var body models.Ticket
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicket(id, body))
}
