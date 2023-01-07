package routes

import (
	// "go-cogoport/middlewares"
	"github.com/gin-gonic/gin"
	controllers "github.com/ChandelShikha/go-cogoport/controllers/api"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.GET("user-list", controllers.UserList)
	v1.POST("create-user", controllers.CreateUser)
	v1.POST("delete-user", controllers.DeleteUser)
	v1.POST("update-user", controllers.UpdateUser)

	return r

}
