package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"github.com/tejas-cogo/go-cogoport/config"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/tickets"
)

func ListTicket(c *gin.Context) {
	var filters models.TicketExtraFilter
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request!")
		return
	}

	ser, db := service.ListTicket(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		data := paginate.New().With(db).Request(c.Request).Response(&ser)
		fmt.Println(data.MaxPage)
		c.JSON(c.Writer.Status(), data)
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
		c.JSON(c.Writer.Status(), "Bad Request")
		return
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
		c.JSON(c.Writer.Status(), "Bad Request")
		return
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
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	if filters.ID <= 0 {
		c.JSON(c.Writer.Status(), "Bad Request!")
		return
	}

	ser, err := service.ListTicketDetail(filters)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		db := config.GetCDB()
		var user models.User
		db.Where("id = ?", ser.TicketReviewer.UserID).First(&user)
		ser.TicketReviewer.User = user
		var role models.AuthRole
		db.Where("id = ?", ser.TicketReviewer.RoleID).First(&role)
		ser.TicketReviewer.Role = role
		c.JSON(c.Writer.Status(), ser)
	}

}

func CreateTicket(c *gin.Context) {
	var body models.Ticket
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.CreateTicket(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {

		c.JSON(c.Writer.Status(), ser)
	}

}

func UpdateTicket(c *gin.Context) {
	var body models.Ticket
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.UpdateTicket(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
