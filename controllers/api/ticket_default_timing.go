package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_timings"
)

// func ListTicketDefaultTiming(c *gin.Context) {
// 	c.JSON(200, service.TicketDefaultTimingList())
// }

func CreateTicketDefaultTiming(c *gin.Context) {
	var ticket_default_timing models.TicketDefaultTiming
	c.BindJSON(&ticket_default_timing)
	c.JSON(200, service.CreateTicketDefaultTiming(ticket_default_timing))
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
	id := body.ID
	c.JSON(200, service.UpdateTicketDefaultTiming(id, body))
}
