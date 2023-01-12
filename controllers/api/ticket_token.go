package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_tokens"
)

func ListTicketToken(c *gin.Context) {
	c.JSON(200, service.ListTicketToken())
}

func CreateTicketToken(c *gin.Context) {
	var ticket_token models.TicketToken
	c.BindJSON(&ticket_token)
	c.JSON(200, service.CreateTicketToken(ticket_token))
}

func CreateTokenTicket(c *gin.Context) {
	token := c.Request.URL.Query().Get("TicketToken")
	var  ticket models.Ticket
	c.BindJSON(&ticket)
	c.JSON(200, service.CreateTokenTicket(token,ticket))
}

func DeleteTicketToken(c *gin.Context) {
	var body models.TicketToken
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketToken(id))
}

// func UpdateTicketToken(c *gin.Context) {
// 	var body models.TicketToken
// 	c.BindJSON(&body)
// 	id := body.ID
// 	c.JSON(200, service.UpdateTicketToken(id, body))
// }