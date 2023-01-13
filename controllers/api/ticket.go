package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

func ListTicket(c *gin.Context) {
	var filters models.Ticket

	filters.Source = c.Request.URL.Query().Get("filters[source]")
	filters.Type = c.Request.URL.Query().Get("filters[type]")
	filters.Priority = c.Request.URL.Query().Get("filters[priority]")
	filters.Status = c.Request.URL.Query().Get("filters[status]")
	tags := c.Request.URL.Query().Get("filters[tags]")
	// c.JSON(200, pg.Response(model, c.Request, &[]Article{}))
	ser, db := service.ListTicket(filters, tags)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	c.BindJSON(&ticket)
	c.JSON(200, service.CreateTicket(ticket))
}

func DeleteTicket(c *gin.Context) {
	var body models.Ticket
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicket(id))
}

func UpdateTicket(c *gin.Context) {
	var body models.Ticket
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicket(id, body))
}
