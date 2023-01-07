package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_spectators"
)

func ListTicketSpectator(c *gin.Context) {
	c.JSON(200, service.ListTicketSpectator())
}

func CreateTicketSpectator(c *gin.Context) {
	var ticket_spectator models.TicketSpectator
	c.BindJSON(&ticket_spectator)
	c.JSON(200, service.CreateTicketSpectator(ticket_spectator))
}

// func DeleteTicketSpectator(c *gin.Context) {
// 	id := c.Request.URL.Query().Get("ID")
// 	c.JSON(200, service.DeleteTicketSpectator(id))
// }

func UpdateTicketSpectator(c *gin.Context) {
	var body models.TicketSpectator
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketSpectator(id, body))
}
