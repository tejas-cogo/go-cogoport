package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"github.com/tejas-cogo/go-cogoport/config"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/tickets"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

func ListTicket(c *gin.Context) {
	var filters models.TicketExtraFilter
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}

	ser, db := service.ListTicket(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		data := paginate.New().With(db).Request(c.Request).Response(&ser)

		items, _ := json.Marshal(data.Items)
		var output []models.TicketData

		// db2 := config.GetCDB()
		err := json.Unmarshal([]byte(items), &output)
		if err != nil {
			print(err)
			c.JSON(400, err)
		}

		var users []string

		for j := 0; j < len(output); j++ {

			// db2.Where("id = ?", output[j].UserID).First(&user)
			// output[j].User = user

			users = append(users, output[j].UserID.String())
		}

		user_data := helpers.GetPartnerUserData(users)

		for j := 0; j < len(user_data); j++ {

			for i := 0; i < len(output); i++ {

				if user_data[i].ID == output[j].User.ID {
					output[i].User.ID = user_data[j].ID
					// output[i].User.Name = user_data[j].Name
					// output[i].User.Email = user_data[j].Email
					// output[i].User.MobileNumber = user_data[j].MobileNumber
					break
				}
			}

		}

		data.Items = output

		c.JSON(c.Writer.Status(), data)
	}
}

func ListTicketTag(c *gin.Context) {
	var Tag string
	Tag = c.Request.URL.Query().Get("Tag")
	ser := service.ListTicketTag(Tag)

	c.JSON(c.Writer.Status(), ser)

}

func GetTicketStats(c *gin.Context) {
	var stats models.TicketStat

	err := c.Bind(&stats)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}

	ser, err := service.GetTicketStats(stats)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func GetTicketGraph(c *gin.Context) {
	var graph models.TicketGraph

	err := c.Bind(&graph)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}

	ser, err := service.GetTicketGraph(graph)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func ListTicketDetail(c *gin.Context) {
	var filters models.TicketExtraFilter

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	if filters.ID <= 0 {
		c.JSON(c.Writer.Status(), "Id required!")
		return
	}

	ser, err := service.ListTicketDetail(filters)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		db := config.GetDB()

		var user models.User
		db.Where("id = ?", ser.TicketReviewer.UserID).First(&user)
		ser.TicketReviewer.User = user

		var t_user models.TicketUser
		fmt.Println(ser.Ticket.UserID, "ser")
		db.Where("system_user_id = ?", ser.Ticket.UserID).First(&t_user)
		ser.TicketUser = t_user

		var role models.AuthRole
		db.Where("id = ?", ser.TicketReviewer.RoleID).First(&role)
		ser.TicketReviewer.Role = role
		c.JSON(c.Writer.Status(), ser)
	}

}

func CreateTicket(c *gin.Context) {
	var body models.Ticket
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	ser, err := service.CreateTicket(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {

		c.JSON(c.Writer.Status(), ser)
	}

}

func UpdateTicket(c *gin.Context) {
	var body models.Ticket
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	ser, err := service.UpdateTicket(body)
	if err != nil {
		c.JSON(c.Writer.Status(), err)
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
