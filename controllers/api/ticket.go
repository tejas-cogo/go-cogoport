package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

func ListTicket(c *gin.Context) {
	var filters models.TicketExtraFilter
	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", err, "status")
		c.JSON(400, "Not Found")
		return
	}

	ser, db, err := service.ListTicket(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func ListTicketTag(c *gin.Context) {
	var Tag string
	Tag = c.Request.URL.Query().Get("Tag")
	ser, err := service.ListTicketTag(Tag)
	if err != nil {
		c.JSON(400,err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func GetTicketStats(c *gin.Context) {
	var stats models.TicketStat

	err := c.Bind(&stats)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	ser, err := service.GetTicketStats(stats)
	if err != nil {
		c.JSON(400,err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func GetTicketGraph(c *gin.Context) {
	var graph models.TicketGraph

	err := c.Bind(&graph)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	ser, err := service.GetTicketGraph(graph)
	if err != nil {
		c.JSON(400,err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func ListTicketDetail(c *gin.Context) {
	var filters models.TicketExtraFilter

	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	c.JSON(200, service.ListTicketDetail(filters))
}

func CreateTicket(c *gin.Context) {
	var body models.Ticket
	c.BindJSON(&body)
	_, err := service.CreateTicket(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), "Successfully created!")
	}

}

func UpdateTicket(c *gin.Context) {
	var body models.Ticket
	c.BindJSON(&body)
	ser, err := service.UpdateTicket(body)
	if err != nil {
		c.JSON(400,err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
