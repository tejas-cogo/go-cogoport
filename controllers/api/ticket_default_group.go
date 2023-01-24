package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_groups"
)

// func ListTicketDefaultGroup(c *gin.Context) {
// 	var ticket_default_group models.TicketDefaultGroup
// 	c.BindJSON(&ticket_default_group)
// 	ser, db := service.ListTicketDefaultGroup(ticket_default_group)
// 	if c.Writer.Status() == 400 {
// 		fmt.Println("status", c.Writer.Status(), "status")
// 		c.JSON(c.Writer.Status(), "Not Found")
// 	} else {
// 		pg := paginate.New()
// 		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
// 	}
// }

func CreateTicketDefaultGroup(c *gin.Context) {
	var ticket_default_group models.TicketDefaultGroup
	c.BindJSON(&ticket_default_group)
	ser, err := service.CreateTicketDefaultGroup(ticket_default_group)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else if ser != "Successfully Created" {
		c.JSON(400, ser)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
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
	c.JSON(200, service.UpdateTicketDefaultGroup(body))
}
