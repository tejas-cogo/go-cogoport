package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func ListTicketActivity(c *gin.Context) {

	var filters models.TicketActivity
	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	ser, db := service.ListTicketActivity(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

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

	ser, mesg, err := service.CreateTicketActivity(filters)

	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else if mesg != "Successfully Created!" {
		c.JSON(400, mesg)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
