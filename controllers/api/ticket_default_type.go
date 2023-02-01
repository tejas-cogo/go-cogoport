package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_types"
)

func ListTicketType(c *gin.Context) {
	var filters models.TicketDefaultType

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Not Found")
	}

	ser := service.ListTicketType(filters)
	c.JSON(c.Writer.Status(), ser)
}

func ListTicketDefaultType(c *gin.Context) {
	var filters models.TicketDefaultFilter

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Not Found")
	}

	ser, db := service.ListTicketDefaultType(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateTicketDefaultType(c *gin.Context) {
	var ticket_default_type models.TicketDefaultType
	c.BindJSON(&ticket_default_type)
	ser, err := service.CreateTicketDefaultType(ticket_default_type)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteTicketDefaultType(c *gin.Context) {
	var body models.TicketDefaultType
	c.BindJSON(&body)
	id := body.ID
	ser, err := service.DeleteTicketDefaultType(id)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateTicketDefaultType(c *gin.Context) {
	var body models.TicketDefaultType
	c.BindJSON(&body)
	ser, err := service.UpdateTicketDefaultType(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
