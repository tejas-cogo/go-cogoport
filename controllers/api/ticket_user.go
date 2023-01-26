package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
)

func ListTicketUser(c *gin.Context) {
	var filters models.TicketUserFilter

	// filters.Name = c.Request.URL.Query().Get("Name")
	// filters.Email = c.Request.URL.Query().Get("Email")
	// filters.MobileNumber = c.Request.URL.Query().Get("MobileNumber")

	// ID := c.Request.URL.Query().Get("SystemUserID")
	// if ID != "" {
	// 	filters.SystemUserID, _ = uuid.Parse(ID)
	// }
	err := c.Bind(&filters)
	if err != nil {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(400, "Not Found")
	}

	ser, db := service.ListTicketUser(filters)
	if c.Writer.Status() == 400 {
		fmt.Println("status", c.Writer.Status(), "status")
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func CreateTicketUser(c *gin.Context) {
	var ticket_user models.TicketUser
	c.BindJSON(&ticket_user)
	c.JSON(200, service.CreateTicketUser(ticket_user))
}

func DeleteTicketUser(c *gin.Context) {
	var body models.TicketUser
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteTicketUser(id))
}

func UpdateTicketUser(c *gin.Context) {
	var body models.TicketUserRole
	c.BindJSON(&body)
	ser, mesg, err := service.UpdateTicketUser(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else if mesg != "Successfully Updated!" {
		c.JSON(400, mesg)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
