package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_groups"
)

func CreateTicketDefaultGroup(c *gin.Context) {
	var ticket_default_group models.TicketDefaultGroup
	c.BindJSON(&ticket_default_group)
	ser, err := service.CreateTicketDefaultGroup(ticket_default_group)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteTicketDefaultGroup(c *gin.Context) {
	var body models.TicketDefaultGroup
	c.BindJSON(&body)
	id := body.ID
	ser, err := service.DeleteTicketDefaultGroup(id)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateTicketDefaultGroup(c *gin.Context) {
	var body models.TicketDefaultGroup
	c.BindJSON(&body)
	ser, err := service.UpdateTicketDefaultGroup(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
