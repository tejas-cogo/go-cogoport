package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_default_roles"
)

func CreateTicketDefaultRole(c *gin.Context) {
	var ticket_default_role models.TicketDefaultRole
	err := c.Bind(&ticket_default_role)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	ser, err := service.CreateTicketDefaultRole(ticket_default_role)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

// func DeleteTicketDefaultRole(c *gin.Context) {
// 	var body models.TicketDefaultRole
// 	err := c.Bind(&body)
// 	if err != nil {
// 		c.JSON(c.Writer.Status(), err.Error())
// 		return
// 	}
// 	id := body.ID
// 	ser, err := service.DeleteTicketDefaultRole(id)
// 	if err != nil {
// 		c.JSON(400, err.Error())
// 	} else {
// 		c.JSON(c.Writer.Status(), ser)
// 	}
// }

// func UpdateTicketDefaultRole(c *gin.Context) {
// 	var body models.TicketDefaultRole
// 	err := c.Bind(&body)
// 	if err != nil {
// 		c.JSON(c.Writer.Status(), err.Error())
// 		return
// 	}
// 	ser, err := service.UpdateTicketDefaultRole(body)
// 	if err != nil {
// 		c.JSON(400, err.Error())
// 	} else {
// 		c.JSON(c.Writer.Status(), ser)
// 	}
// }
