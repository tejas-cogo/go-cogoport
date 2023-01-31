package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_spectators"
)

func ListTicketSpectator(c *gin.Context) {
	var filters models.TicketSpectator

	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	ser, db, err := service.ListTicketSpectator(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateTicketSpectator(c *gin.Context) {
	var ticket_spectator models.TicketSpectator
	c.BindJSON(&ticket_spectator)
	ser, err := service.CreateTicketSpectator(ticket_spectator)
	if err != nil {
		c.JSON(400,err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteTicketSpectator(c *gin.Context) {
	var body models.TicketSpectator
	c.BindJSON(&body)
	id := body.ID
	ser, err := service.DeleteTicketSpectator(id)
	if err != nil {
		c.JSON(400,err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
