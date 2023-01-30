package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_tokens"
)

func CreateTicketToken(c *gin.Context) {
	var body models.TicketUser
	c.BindJSON(&body)
	c.JSON(200, service.CreateTicketToken(body))
}

func CreateTokenTicket(c *gin.Context) {

	var token_ticket models.TicketToken
	c.BindJSON(&token_ticket)
	c.JSON(200, service.CreateTokenTicket(token_ticket))
}

func DeleteTicketToken(c *gin.Context) {
	var body models.TicketToken
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketToken(id))
}
