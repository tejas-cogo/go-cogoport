package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_groups"
)

// func ListTicketDefaultGroup(c *gin.Context) {
// 	c.JSON(200, service.ListTicketDefaultGroup())
// }

func CreateTicketDefaultGroup(c *gin.Context) {
	var ticket_default_group models.TicketDefaultGroup
	c.BindJSON(&ticket_default_group)
	c.JSON(200, service.CreateTicketDefaultGroup(ticket_default_group))
}

func DeleteTicketDefaultGroup(c *gin.Context) {
	var body models.TicketDefaultGroup
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketDefaultGroup(id))
}

func UpdateTicketDefaultGroup(c *gin.Context) {
	var body models.TicketDefaultGroup
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketDefaultGroup(id, body))
}
