package routes

import (
	// "go-cogoport/middlewares"
	"github.com/gin-gonic/gin"
	controllers "github.com/tejas-cogo/go-cogoport/controllers/api"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.POST("user-list", controllers.UserList)
	return r

}
