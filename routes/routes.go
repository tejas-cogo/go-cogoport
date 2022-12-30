package routes

import (
	// "go-cogoport/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/tejas-cogo/go-cogogoport/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.POST("user-list", controller.UserList)
	return r

}
