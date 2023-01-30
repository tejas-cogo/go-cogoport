package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
)

// func ListTicketReviewer(c *gin.Context) {
// 	var filters models.TicketReviewer

// 	// TicketID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_id]"))
// 	// filters.TicketID = uint(TicketID)

// 	err := c.Bind(&filters)
// 	if err != nil {
// 		fmt.Println("status", c.Writer.Status(), "status")
// 		c.JSON(400, "Not Found")
// 	}

// 	// TicketUserID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[ticket_user_id]"))
// 	// filters.TicketUserID = uint(TicketUserID)

// 	ser, db := service.ListTicketReviewer(filters)
// 	if c.Writer.Status() == 400 {
// 		fmt.Println("status", c.Writer.Status(), "status")
// 		c.JSON(c.Writer.Status(), "Not Found")
// 	} else {
// 		pg := paginate.New()
// 		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
// 	}
// }


func ReassignTicketReviewer(c *gin.Context) {
	var body models.ReviewerActivity
	c.BindJSON(&body)
	c.JSON(200, service.ReassignTicketReviewer(body))
}
