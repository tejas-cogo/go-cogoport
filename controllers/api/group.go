package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/groups"
)

func ListGroup(c *gin.Context) {
	var filters models.Group
	filters.Status = c.Request.URL.Query().Get("filters[status]")
	tags := c.Request.URL.Query().Get("filters[tags]")
	ser, db := service.ListGroup(filters, tags)
	pg := paginate.New()
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func CreateGroup(c *gin.Context) {
	var group models.Group
	c.BindJSON(&group)
	c.JSON(200, service.CreateGroup(group))
}

func DeleteGroup(c *gin.Context) {
	var body models.Group
	c.BindJSON(&body)
	id := body.ID
	c.JSON(200, service.DeleteGroup(id))
}

func UpdateGroup(c *gin.Context) {
	var body models.Group
	c.BindJSON(&body)
	c.JSON(200, service.UpdateGroup(body))
}
