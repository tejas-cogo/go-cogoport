package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

func ListTicket(c *gin.Context) {
	var filters models.Ticket

	ID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[id]"))
	filters.ID = uint(ID)

	filters.Source = c.Request.URL.Query().Get("filters[source]")
	filters.Type = c.Request.URL.Query().Get("filters[type]")
	filters.Priority = c.Request.URL.Query().Get("filters[priority]")
	filters.Status = c.Request.URL.Query().Get("filters[status]")
	// filters.Tags[0] = c.Request.URL.Query().Get("filters[tags]")
	// c.JSON(200, pg.Response(model, c.Request, &[]Article{}))
	ser, db := service.ListTicket(filters)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func ListTicketDetail(c *gin.Context) {
	var filters models.Ticket

	ID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[id]"))
	filters.ID = uint(ID)

	// TicketSpectatorID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_spectator_id]"))
	// filters.TicketSpectatorID = uint(TicketSpectatorID)

	// TicketReviewerID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_reviewer_id]"))
	// filters.TicketReviewerID = uint(TicketReviewerID)

	// TicketActivityID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_activity_id]"))
	// filters.TicketActivityID = uint(TicketActivityID)

	ser, db := service.ListTicketDetail(filters)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
	// c.JSON(200, service.ListTicketDetail(filters))
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
