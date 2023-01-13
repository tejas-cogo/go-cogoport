package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func ListTicketActivity(c *gin.Context) {
	var filters models.TicketActivity

	TicketID, _ :=strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_id]"))
	filters.TicketID = uint(TicketID)

	filters.UserType = c.Request.URL.Query().Get("filters[user_type]")

	TicketUserID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_user_id]"))
	filters.TicketUserID = uint(TicketUserID)

	filters.IsRead,_ = strconv.ParseBool(c.Request.URL.Query().Get("filters[is_read]"))

	ser, db := service.ListTicketActivity(filters)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func CreateTicketActivity(c *gin.Context) {
	var ticket_activity models.TicketActivity
	c.BindJSON(&ticket_activity)
	c.JSON(200, service.CreateTicketActivity(ticket_activity))
}

func DeleteTicketActivity(c *gin.Context) {
	var body models.TicketActivity
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketActivity(id))
}

func UpdateTicketActivity(c *gin.Context) {
	var body models.TicketActivity
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketActivity(id, body))
}
