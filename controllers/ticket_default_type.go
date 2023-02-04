package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	group_service "github.com/tejas-cogo/go-cogoport/services/api/ticket_default_groups"
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
		data := paginate.New().Response(db, c.Request, &ser)
		items, _ := json.Marshal(data.Items)
		var output []models.TicketDefault

		err := json.Unmarshal([]byte(items), &output)
		if err != nil {
			print(err)
		}

		list := make([]interface{}, 0)
		for _, value := range output {
			fmt.Println(value)
			var f models.TicketDefaultGroup
			f.TicketDefaultTypeID = value.ID

			value.TicketDefaultGroupTypeQuery, _ = group_service.ListTicketDefaultGroup(f)
			list = append(list, value)

		}

		data.Items = list

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
