package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"github.com/tejas-cogo/go-cogoport/config"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_tokens"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

func ListTokenTicketDetail(c *gin.Context) {
	var filters models.TokenFilter
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	if filters.TicketToken == "" {
		c.JSON(c.Writer.Status(), "Token Required!")
		return
	}
	ser, err := service.ListTokenTicketDetail(filters)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		db := config.GetDB()

		var user models.User
		db.Where("id = ?", ser.TicketReviewer.UserID).First(&user)
		ser.TicketReviewer.User = user
		var ticket_user models.TicketUser
		db.Where("system_user_id = ?", ser.Ticket.UserID).First(&ticket_user)
		ser.TicketUser = ticket_user

		var role models.AuthRole
		db.Where("id = ?", ser.TicketReviewer.RoleID).First(&role)
		ser.TicketReviewer.Role = role
		c.JSON(c.Writer.Status(), ser)
	}
}

func ListTokenTicketActivity(c *gin.Context) {
	var filters models.TokenFilter
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	if filters.TicketToken == "" {
		c.JSON(c.Writer.Status(), "Token Required!")
		return
	}
	ser, db, err := service.ListTokenTicketActivity(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), err)
	} else {
		data := paginate.New().With(db).Request(c.Request).Response(&ser)
		items, _ := json.Marshal(data.Items)
		var output []models.TicketActivityData

		db2 := config.GetDB()
		err := json.Unmarshal([]byte(items), &output)
		if err != nil {
			print(err)
			c.JSON(400, err)
		}
		var users []string

		for j := 0; j < len(output); j++ {
			if output[j].UserType == "ticket_user" {
				var user models.TicketUser
				db2.Where("id = ?", output[j].Ticket.TicketUserID).First(&user)
				output[j].TicketUser = user
			} else {

				users = append(users, output[j].UserID.String())

			}
		}

		user_data := helpers.GetUserData(users)

		for j := 0; j < len(output); j++ {
			if output[j].UserType != "ticket_user" {
				for i := 0; i < len(user_data); i++ {
					if user_data[i].ID == output[j].UserID {
						output[j].TicketUser.SystemUserID = user_data[i].ID
						output[j].TicketUser.Name = user_data[i].Name
						output[j].TicketUser.Email = user_data[i].Email
						output[j].TicketUser.MobileNumber = user_data[i].MobileNumber
						break
					}
				}
			}
		}
		data.Items = output
		c.JSON(c.Writer.Status(), data)
	}
}

func GetTicketToken(c *gin.Context) {
	var body models.TicketUser
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	ser, err := service.GetTicketToken(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func CreateTokenTicket(c *gin.Context) {
	var token_filter models.TokenFilter
	err := c.Bind(&token_filter)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	if token_filter.TicketToken == "" {
		c.JSON(c.Writer.Status(), "Token Required!")
		return
	}
	ser, err := service.CreateTokenTicket(token_filter)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func CreateTokenTicketActivity(c *gin.Context) {

	var token_filter models.TokenActivity
	err := c.Bind(&token_filter)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	if token_filter.TicketToken == "" {
		c.JSON(c.Writer.Status(), "Token Required!")
		return
	}
	ser, err := service.CreateTokenTicketActivity(token_filter)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

// func DeleteTicketToken(c *gin.Context) {
// 	var body models.TicketToken
// 	c.BindJSON(&body)
// 	id := body.ID
// 	ser, err := service.DeleteTicketToken(id)
// 	if err != nil {
// 		c.JSON(c.Writer.Status(), err)
// 	} else {
// 		c.JSON(c.Writer.Status(), ser)
// 	}
// }

func UpdateTokenTicket(c *gin.Context) {
	var body models.TokenFilter
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	if body.TicketToken == "" {
		c.JSON(c.Writer.Status(), "Token Required!")
		return
	}
	ser, err := service.UpdateTokenTicket(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func ListTokenTicketType(c *gin.Context) {
	var body models.TokenFilter

	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	if body.TicketToken == "" {
		c.JSON(c.Writer.Status(), "Token Required!")
		return
	}

	ser, db := service.ListTokenTicketType(body)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}
