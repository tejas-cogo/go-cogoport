package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_tokens"
)

func ListTokenTicketDetail(c *gin.Context) {
	var filters models.TokenFilter
	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	c.JSON(200, service.ListTokenTicketDetail(filters))
}

func CreateTicketToken(c *gin.Context) {
	var body models.TicketUser
	c.BindJSON(&body)
	c.JSON(200, service.CreateTicketToken(body))
}

func CreateTokenTicket(c *gin.Context) {

	var token_filter models.TokenFilter
	c.BindJSON(&token_filter)
	c.JSON(200, service.CreateTokenTicket(token_filter))
}

func DeleteTicketToken(c *gin.Context) {
	var body models.TicketToken
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketToken(id))
}
