package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_tokens"
)

func ListTokenTicketDetail(c *gin.Context) {
	var filters models.TokenFilter
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(400, "Not Found")
	}

	c.JSON(200, service.ListTokenTicketDetail(filters))
}

func CreateTicketToken(c *gin.Context) {
	var body models.TicketUser
	c.BindJSON(&body)
	ser, err := service.CreateTicketToken(body)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func CreateTokenTicket(c *gin.Context) {

	var token_filter models.TokenFilter
	c.BindJSON(&token_filter)
	ser, err := service.CreateTokenTicket(token_filter)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteTicketToken(c *gin.Context) {
	var body models.TicketToken
	c.BindJSON(&body)
	id := body.ID
	ser, err := service.DeleteTicketToken(id)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateTokenTicket(c *gin.Context) {
	var body models.TokenFilter
	c.BindJSON(&body)
	ser, err := service.UpdateTokenTicket(body)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
