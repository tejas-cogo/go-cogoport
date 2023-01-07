package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_types"
)

// func ListTicketDefaultType(c *gin.Context) {
// 	c.JSON(200, service.TicketDefaultTypeList())
// }

// func CreateTicketDefaultType(c *gin.Context) {
// 	var ticket_default_type models.TicketDefaultType
// 	c.BindJSON(&ticket_default_type)
// 	c.JSON(200, service.CreateTicketDefaultType(ticket_default_type))
// }

// func DeleteTicketDefaultType(c *gin.Context) {
// 	id := c.Request.URL.Query().Get("ID")
// 	c.JSON(200, service.DeleteTicketDefaultType(id))
// }

func UpdateTicketDefaultType(c *gin.Context) {
	var body models.TicketDefaultType
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.UpdateTicketDefaultType(id, body))
}
