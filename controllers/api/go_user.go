package controllers

import (
	"github.com/gin-gonic/gin"
	hlp "github.com/tejas-cogo/go-cogoport/helpers"
	service "github.com/tejas-cogo/go-cogoport/services/api/users"
)

func UserList(c *gin.Context) {
	resp := service.UserList()
	hlp.Respond(c.Writer, resp)
}
