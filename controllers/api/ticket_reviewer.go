package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
)

func ListTicketReviewer(c *gin.Context) {
	var filters models.TicketReviewer

	TicketID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_id]"))
	filters.TicketID = uint(TicketID)

	TicketUserID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_user_id]"))
	filters.TicketUserID = uint(TicketUserID)

	ser, db := service.ListTicketReviewer(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

// func CreateTicketReviewer(c *gin.Context) {
// 	var body models.Filter
// 	c.BindJSON(&body)
// 	c.JSON(200, service.CreateTicketReviewer(body))
// }

func ReassignTicketReviewer(c *gin.Context) {
	var body models.ReviewerActivity
	c.BindJSON(&body)
	c.JSON(200, service.ReassignTicketReviewer(body))
}

// func DeleteTicketReviewer(c *gin.Context) {
// 	var body models.TicketReviewer
// 	c.BindJSON(&body)
// 	id := body.ID
// 	c.JSON(200, service.DeleteTicketReviewer(id))
// }

// func UpdateTicketReviewer(c *gin.Context) {
// 	var body models.TicketReviewer
// 	c.BindJSON(&body)
// 	c.JSON(200, service.UpdateTicketReviewer(body))
// }
