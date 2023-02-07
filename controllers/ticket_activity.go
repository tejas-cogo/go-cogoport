package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"github.com/tejas-cogo/go-cogoport/config"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_activities"
)

func CreateTicketActivity(c *gin.Context) {
	var body models.Activity
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	var filters models.Filter
	filters.Activity.TicketID = body.TicketID
	filters.TicketActivity.UserID = body.PerformedByID
	filters.TicketActivity.Type = body.Type
	filters.TicketActivity.UserType = body.UserType
	filters.TicketActivity.Description = body.Description
	filters.TicketActivity.Data = body.Data
	filters.TicketActivity.Status = body.Status

	ser, err := service.CreateTicketActivity(filters)

	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func ListTicketActivity(c *gin.Context) {
	var filters models.TicketActivity

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}

	ser, db, err := service.ListTicketActivity(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		data := pg.Response(db, c.Request, &ser)
		items, _ := json.Marshal(data.Items)
		var output []models.TicketActivityData

		db2 := config.GetCDB()
		err := json.Unmarshal([]byte(items), &output)
		if err != nil {
			print(err)
			c.JSON(400, err)
		}

		for j := 0; j < len(output); j++ {
			if output[j].UserType != "ticket_user" {
				var user models.TicketUser
				db2.Where("id = ?", output[j].Ticket.TicketUserID).First(&user)
				fmt.Println("output[j].UserID", output[j].UserID)
				output[j].TicketUser = user
			} else {
				fmt.Println("output[j].UserID", output[j].UserID)

				var user models.User
				db2.Where("system_user_id = ?", output[j].UserID).First(&user)
				output[j].TicketUser.SystemUserID = user.ID
				output[j].TicketUser.Name = user.Name
				output[j].TicketUser.Email = user.Email
				output[j].TicketUser.MobileNumber = user.MobileNumber
			}

		}
		data.Items = output
		c.JSON(c.Writer.Status(), data)
	}
}
