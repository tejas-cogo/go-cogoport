package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_tokens"
)

func ListTokenTicketDetail(c *gin.Context) {
	var filters models.TokenFilter
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.ListTokenTicketDetail(filters)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func ListTokenTicketActivity(c *gin.Context) {
	var filters models.TokenFilter
	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, db, err := service.ListTokenTicketActivity(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), err)
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func GetTicketToken(c *gin.Context) {
	var body models.TicketUser
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), "Bad Request")
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
		c.JSON(c.Writer.Status(), "Bad Request")
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
		c.JSON(c.Writer.Status(), "Bad Request")
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
		c.JSON(c.Writer.Status(), "Bad Request")
		return
	}
	ser, err := service.UpdateTokenTicket(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
