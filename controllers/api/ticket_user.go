package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
)

func ListTicketUser(c *gin.Context) {
	var filters models.TicketUserFilter

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(400, "Not Found")
	}

	ser, db, err := service.ListTicketUser(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateTicketUser(c *gin.Context) {
	var ticket_user models.TicketUser
	c.BindJSON(&ticket_user)
	ser, err := service.CreateTicketUser(ticket_user)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateTicketUser(c *gin.Context) {
	var body models.TicketUserRole
	c.BindJSON(&body)
	ser, err := service.UpdateTicketUser(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
