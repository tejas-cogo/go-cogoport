package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
)

// func CreateTicketAudit(c *gin.Context) {
// 	var ticket_audit models.TicketAudit
// 	c.BindJSON(&ticket_audit)
// 	c.JSON(200, service.CreateTicketAudit(ticket_audit))
// }

func DeleteTicketAudit(c *gin.Context) {
	var body models.TicketAudit
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketAudit(id))
}

func UpdateTicketAudit(c *gin.Context) {
	var body models.TicketAudit
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketAudit(id, body))
}
