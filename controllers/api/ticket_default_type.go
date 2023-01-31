package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_types"
)

func ListTicketType(c *gin.Context) {
	var filters models.TicketDefaultType

	// filters.TicketType = c.Request.URL.Query().Get("filters[ticket_type]")

	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	ser, db := service.ListTicketType(filters)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func ListTicketDefaultType(c *gin.Context) {
	fmt.Println("filterssssssss")
	var filters models.TicketDefaultFilter

	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}
	fmt.Println("filters", filters)

	ser, db := service.ListTicketDefaultType(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
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
	} else if ser != "Successfully Created!" {
		c.JSON(400, ser)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteTicketDefaultType(c *gin.Context) {
	var body models.TicketDefaultType
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketDefaultType(id))
}

func UpdateTicketDefaultType(c *gin.Context) {
	var body models.TicketDefaultType
	c.BindJSON(&body)
	c.JSON(200, service.UpdateTicketDefaultType(body))
}
