package controllers

import(
	"github.com/gin-gonic/gin"
	user_service "github.com/tejas-cogo/go-cogoport/services/api/users"

)


func UserList(c *gin.Context) {
	var user_service service.UserService
	resp := service.UserList()
	u.Respond(c.Writer, resp)
}
