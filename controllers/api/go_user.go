package controllers

import (
	"github.com/gin-gonic/gin"
	service "github.com/tejas-cogo/go-cogoport/services/api/users"
)

func UserList(c *gin.Context) {
	c.JSON(200, service.UserList())
}
