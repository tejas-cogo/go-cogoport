package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
)

func ListTicketUser(c *gin.Context) {
	var filters models.TicketUser

	filters.Name = c.Request.URL.Query().Get("filters[name]")

	ser, db := service.ListTicketUser(filters)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func CreateTicketUser(c *gin.Context) {
	var ticket_user models.TicketUser
	c.BindJSON(&ticket_user)
	c.JSON(200, service.CreateTicketUser(ticket_user))
}

func DeleteTicketUser(c *gin.Context) {
	var body models.TicketUser
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketUser(id))
}

func UpdateTicketUser(c *gin.Context) {
	var body models.TicketUser
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketUser(id, body))
}
