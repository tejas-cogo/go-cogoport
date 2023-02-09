package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_users"
)

func CreateTicketUser(c *gin.Context) {
	var ticket_user models.TicketUser
	err := c.Bind(&ticket_user)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	ser, err := service.CreateTicketUser(ticket_user)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

