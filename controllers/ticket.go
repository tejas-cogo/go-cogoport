package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/tickets"
)

func ListTicket(c *gin.Context) {
	var filters models.TicketExtraFilter
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Not Found")
		return
	}

	ser, db := service.ListTicket(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func ListTicketTag(c *gin.Context) {
	var Tag string
	Tag = c.Request.URL.Query().Get("Tag")
	ser := service.ListTicketTag(Tag)

	c.JSON(c.Writer.Status(), ser)

}

func GetTicketStats(c *gin.Context) {
	var stats models.TicketStat

	err := c.Bind(&stats)
	if err != nil {
		c.JSON(c.Writer.Status(), "Not Found")
	}

	ser, err := service.GetTicketStats(stats)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func GetTicketGraph(c *gin.Context) {
	var graph models.TicketGraph

	err := c.Bind(&graph)
	if err != nil {
		c.JSON(c.Writer.Status(), "Not Found")
	}

	ser, err := service.GetTicketGraph(graph)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func ListTicketDetail(c *gin.Context) {
	var filters models.TicketExtraFilter

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Not Found")
	}

	c.JSON(c.Writer.Status(), service.ListTicketDetail(filters))
}

func CreateTicket(c *gin.Context) {
	var body models.Ticket
	c.BindJSON(&body)
	ser, err := service.CreateTicket(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}

}

func UpdateTicket(c *gin.Context) {
	var body models.Ticket
	c.BindJSON(&body)
	ser, err := service.UpdateTicket(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
