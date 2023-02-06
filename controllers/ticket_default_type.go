package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"github.com/tejas-cogo/go-cogoport/config"
	models "github.com/tejas-cogo/go-cogoport/models"

	role_service "github.com/tejas-cogo/go-cogoport/services/api/ticket_default_roles"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_default_types"
)

func ListTicketType(c *gin.Context) {
	var filters models.TicketDefaultType

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}

	ser, db := service.ListTicketType(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func ListTicketDefaultType(c *gin.Context) {
	var filters models.TicketDefaultFilter

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}

	ser, db := service.ListTicketDefaultType(filters)

	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		data := paginate.New().With(db).Request(c.Request).Response(&ser)
		items, _ := json.Marshal(data.Items)
		var output []models.TicketDefault

		err := json.Unmarshal([]byte(items), &output)
		if err != nil {
			print(err)
			c.JSON(400, err)
		}

		db2 := config.GetDB()

		for j := 0; j < len(output); j++ {
			var f models.TicketDefaultRole
			f.TicketDefaultTypeID = output[j].ID
			output[j].TicketDefaultRole, _ = role_service.ListTicketDefaultRole(f)

			for i := 0; i < len(output[j].TicketDefaultRole); i++ {
				var user models.User
				users := db2.Where("id = ?", output[j].TicketDefaultRole[i].UserID).First(&user)
				if users.RowsAffected > 0 {
					output[j].TicketDefaultRole[i].User = user
				}

				var auth_role models.AuthRole
				auth_roles := db2.Where("id = ?", output[j].TicketDefaultRole[i].RoleID).First(&auth_role)
				if auth_roles.RowsAffected > 0 {
					fmt.Println("auth_role", auth_role)
					output[j].TicketDefaultRole[i].Role = auth_role
				}
			}
			fmt.Println(output[j].TicketDefaultRole)

		}

		data.Items = output

		c.JSON(c.Writer.Status(), data)
	}
}

func CreateTicketDefaultType(c *gin.Context) {
	var ticket_default_type models.TicketDefaultType
	err := c.Bind(&ticket_default_type)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.CreateTicketDefaultType(ticket_default_type)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteTicketDefaultType(c *gin.Context) {
	var body models.TicketDefaultType
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	id := body.ID
	ser, err := service.DeleteTicketDefaultType(id)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateTicketDefaultType(c *gin.Context) {
	var body models.TicketDefaultType
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.UpdateTicketDefaultType(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
