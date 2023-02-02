package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_default_timings"
)

func CreateTicketDefaultTiming(c *gin.Context) {
	var ticket_default_timing models.TicketDefaultTiming
	c.BindJSON(&ticket_default_timing)

	ser, err := service.CreateTicketDefaultTiming(ticket_default_timing)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteTicketDefaultTiming(c *gin.Context) {
	var body models.TicketDefaultTiming
	c.BindJSON(&body)
	id := body.ID
	ser, err := service.DeleteTicketDefaultTiming(id)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateTicketDefaultTiming(c *gin.Context) {
	var body models.TicketDefaultTiming
	c.BindJSON(&body)
	ser, err := service.UpdateTicketDefaultTiming(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
