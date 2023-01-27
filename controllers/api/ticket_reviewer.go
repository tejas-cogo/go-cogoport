package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
)

func ListTicketReviewer(c *gin.Context) {
	var filters models.TicketReviewer

	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	ser, db := service.ListTicketReviewer(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}


func ReassignTicketReviewer(c *gin.Context) {
	var body models.ReviewerActivity
	c.BindJSON(&body)
	c.JSON(200, service.ReassignTicketReviewer(body))
}
