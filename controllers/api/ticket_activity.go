package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func ListTicketActivity(c *gin.Context) {

	var filters models.TicketActivity
	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	// TicketID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_id]"))
	// filters.TicketID = uint(TicketID)

	// filters.UserType = c.Request.URL.Query().Get("filters[user_type]")

	// TicketUserID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_user_id]"))
	// filters.TicketUserID = uint(TicketUserID)

	// filters.IsRead, _ = strconv.ParseBool(c.Request.URL.Query().Get("filters[is_read]"))

	ser, db := service.ListTicketActivity(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateTicketActivity(c *gin.Context) {
	var body models.Activity
	c.BindJSON(&body)
	var filters models.Filter
	filters.Activity.TicketID = body.TicketID
	filters.Activity.PerformedByID = body.PerformedByID
	filters.TicketActivity.Type = body.Type
	filters.TicketActivity.Description = body.Description
	filters.TicketActivity.Data = body.Data
	filters.TicketActivity.Status = body.Status

	c.JSON(200, service.CreateTicketActivity(filters))
}

// func DeleteTicketActivity(c *gin.Context) {
// 	var body models.TicketActivity
// 	c.BindJSON(&body)
// 	id := body.ID
// 	c.JSON(200, service.DeleteTicketActivity(id))
// }

// func UpdateTicketActivity(c *gin.Context) {
// 	var body models.TicketActivity
// 	c.BindJSON(&body)
// 	c.JSON(200, service.UpdateTicketActivity(body))
// }
