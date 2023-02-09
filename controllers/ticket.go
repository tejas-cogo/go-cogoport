package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
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

		err := json.Unmarshal([]byte(items), &output)
		if err != nil {
			print(err)
			c.JSON(400, err)
		}

		var users []string

		for j := 0; j < len(output); j++ {
			users = append(users, output[j].UserID.String())
		}

		user_data := helpers.GetUserData(users)

		for j := 0; j < len(output); j++ {
			for i := 0; i < len(user_data); i++ {

				if user_data[i].ID == output[j].UserID {
					output[j].User.ID = user_data[i].ID
					output[j].User.Name = user_data[i].Name
					output[j].User.Email = user_data[i].Email
					output[j].User.MobileNumber = user_data[i].MobileNumber
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
