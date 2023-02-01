package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_reviewers"
)

func ReassignTicketReviewer(c *gin.Context) {
	var body models.ReviewerActivity
	c.BindJSON(&body)
	ser, err := service.ReassignTicketReviewer(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
