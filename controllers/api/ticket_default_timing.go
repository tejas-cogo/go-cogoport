package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_timings"
)

func ListTicketDefaultTiming(c *gin.Context) {
	var filters models.TicketDefaultTiming
	filters.TicketPriority = c.Request.URL.Query().Get("filters[ticket_priority]")
	ser, db := service.ListTicketDefaultTiming(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

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
	c.JSON(200, service.UpdateTicketDefaultTiming(body))
}
