package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_tasks"
)

func ListTicketTask(c *gin.Context) {
	c.JSON(200, service.ListTicketTask())
}

func CreateTicketTask(c *gin.Context) {
	var ticket_task models.TicketTask
	c.BindJSON(&ticket_task)
	c.JSON(200, service.CreateTicketTask(ticket_task))
}

func DeleteTicketTask(c *gin.Context) {
	var body models.TicketTask
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketTask(id))
}

func UpdateTicketTask(c *gin.Context) {
	var body models.TicketTask
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketTask(id, body))
}
