package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func CreateTicketActivity(c *gin.Context) {
	var body models.Activity
	c.BindJSON(&body)
	var filters models.Filter
	filters.Activity.TicketID = body.TicketID
	filters.Activity.PerformedByID = body.PerformedByID
	filters.TicketActivity.Type = body.Type
	filters.TicketActivity.Description = body.Description
	filters.TicketActivity.Data = body.Data
	filters.TicketActivity.Status = body.Status

	ser, err := service.CreateTicketActivity(filters)

	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
