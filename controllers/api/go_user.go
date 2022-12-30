package controllers

import (
	"github.com/gin-gonic/gin"
	user_service "github.com/tejas-cogo/go-cogoport/services/api/users"
)

func UserList(c *gin.Context) {
	// var user service.UserService
	resp := user_service.UserList()
	// u.Respond(c.Writer, resp)
	return resp
}
