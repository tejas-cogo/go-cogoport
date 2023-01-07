package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
)

func ListTicketReviewer(c *gin.Context) {
	c.JSON(200, service.TicketReviewerList())
}

func CreateTicketReviewer(c *gin.Context) {
	var ticket_reviewer models.TicketReviewer
	c.BindJSON(&ticket_reviewer)
	c.JSON(200, service.CreateTicketReviewer(ticket_reviewer))
}

func DeleteTicketReviewer(c *gin.Context) {
	id := c.Request.URL.Query().Get("ID")
	c.JSON(200, service.DeleteTicketReviewer(id))
}

func UpdateTicketReviewer(c *gin.Context) {
	var body models.TicketReviewer
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketReviewer(id, body))
}
