package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_types"
)

func ListTicketDefaultType(c *gin.Context) {
	var filters models.TicketDefaultType

	filters.TicketType = c.Request.URL.Query().Get("filters[ticket_type]")

	ser, db := service.ListTicketDefaultType(filters)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func ListTicketDefault(c *gin.Context) {
	var filters models.Filter

	filters.TicketDefaultType.TicketType = c.Request.URL.Query().Get("filters[ticket_type]")

	ser, db := service.ListTicketDefault(filters)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func CreateTicketDefaultType(c *gin.Context) {
	var ticket_default_type models.TicketDefaultType
	c.BindJSON(&ticket_default_type)
	c.JSON(200, service.CreateTicketDefaultType(ticket_default_type))
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
