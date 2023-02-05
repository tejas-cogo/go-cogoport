package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_users"
)

// func ListTicketUser(c *gin.Context) {
// 	var filters models.TicketUserFilter

// 	err := c.Bind(&filters)
// 	if err != nil {
// 		c.JSON(c.Writer.Status(), "Bad Request")
// 		return
// 	}

// 	ser, db, err := service.ListTicketUser(filters)
// 	if c.Writer.Status() == 400 {
// 		c.JSON(c.Writer.Status(), "Not Found")
// 	} else {
// 		pg := paginate.New()
// 		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
// 	}
// }

func CreateTicketUser(c *gin.Context) {
	var ticket_user models.TicketUser
	err := c.Bind(&ticket_user)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.CreateTicketUser(ticket_user)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

// func UpdateTicketUser(c *gin.Context) {
// 	var body models.TicketUser
// 	err := c.Bind(&body)
// 	if err != nil {
// 		c.JSON(c.Writer.Status(), "Bad Request")
// 		return
// 	}
// 	ser, err := service.UpdateTicketUser(body)
// 	if err != nil {
// 		c.JSON(400, err.Error())
// 	} else {
// 		c.JSON(c.Writer.Status(), ser)
// 	}
// }
