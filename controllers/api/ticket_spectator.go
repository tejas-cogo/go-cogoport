package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_spectators"
)

func ListTicketSpectator(c *gin.Context) {
	var filters models.TicketSpectator

	// TicketID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_id]"))
	// filters.TicketID = uint(TicketID)

	// TicketUserID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_user_id]"))
	// filters.TicketUserID = uint(TicketUserID)

	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	ser, db := service.ListTicketSpectator(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateTicketSpectator(c *gin.Context) {
	var ticket_spectator models.TicketSpectator
	c.BindJSON(&ticket_spectator)
	c.JSON(200, service.CreateTicketSpectator(ticket_spectator))
}

func DeleteTicketSpectator(c *gin.Context) {
	var body models.TicketSpectator
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketSpectator(id))
}

// func UpdateTicketSpectator(c *gin.Context) {
// 	var body models.TicketSpectator
// 	c.BindJSON(&body)
// 	id := body.ID
// 	c.JSON(200, service.UpdateTicketSpectator(id, body))
// }
