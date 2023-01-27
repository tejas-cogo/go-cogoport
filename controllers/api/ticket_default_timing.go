package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_timings"
)

func CreateTicketDefaultTiming(c *gin.Context) {
	var ticket_default_timing models.TicketDefaultTiming
	c.BindJSON(&ticket_default_timing)

	ser, err := service.CreateTicketDefaultTiming(ticket_default_timing)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else if ser != "Successfully Created!" {
		c.JSON(400, ser)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteTicketDefaultTiming(c *gin.Context) {
	var body models.TicketDefaultTiming
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketDefaultTiming(id))
}

func UpdateTicketDefaultTiming(c *gin.Context) {
	var body models.TicketDefaultTiming
	c.BindJSON(&body)
	c.JSON(200, service.UpdateTicketDefaultTiming(body))
}
