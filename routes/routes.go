package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/tejas-cogo/go-cogoport/controllers"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())

	ticket_system := r.Group("")

	ticket_system.GET("list_groups", controllers.ListGroup)
	ticket_system.POST("create_group", controllers.CreateGroup)
	ticket_system.DELETE("delete_group", controllers.DeleteGroup)
	ticket_system.PUT("update_group", controllers.UpdateGroup)
	ticket_system.GET("list_group_tags", controllers.ListGroupTag)

	ticket_system.GET("list_group_members", controllers.ListGroupMember)
	ticket_system.POST("create_group_member", controllers.CreateGroupMember)
	ticket_system.DELETE("delete_group_member", controllers.DeleteGroupMember)
	ticket_system.PUT("update_group_member", controllers.UpdateGroupMember)

	ticket_system.GET("list_roles", controllers.ListRole)
	ticket_system.POST("create_role", controllers.CreateRole)
	ticket_system.DELETE("delete_role", controllers.DeleteRole)
	ticket_system.PUT("update_role", controllers.UpdateRole)

	ticket_system.GET("list_ticket_users", controllers.ListTicketUser)
	// ticket_system.POST("create_ticket_user", controllers.CreateTicketUser)
	ticket_system.PUT("update_ticket_user", controllers.UpdateTicketUser)

	ticket_system.GET("list_ticket_details", controllers.ListTicketDetail)
	ticket_system.GET("list_tickets", controllers.ListTicket)
	ticket_system.GET("list_ticket_tags", controllers.ListTicketTag)
	ticket_system.GET("get_ticket_stats", controllers.GetTicketStats)
	ticket_system.GET("get_ticket_graph", controllers.GetTicketGraph)
	ticket_system.POST("create_ticket", controllers.CreateTicket)
	ticket_system.PUT("update_ticket", controllers.UpdateTicket)

	ticket_system.POST("create_ticket_activity", controllers.CreateTicketActivity)
	ticket_system.GET("list_ticket_activities", controllers.ListTicketActivity)

	ticket_system.GET("list_ticket_spectators", controllers.ListTicketSpectator)
	ticket_system.POST("create_ticket_spectator", controllers.CreateTicketSpectator)
	ticket_system.DELETE("delete_ticket_spectator", controllers.DeleteTicketSpectator)

	ticket_system.POST("reassign_ticket_reviewer", controllers.ReassignTicketReviewer)

	ticket_system.POST("create_ticket_default_group", controllers.CreateTicketDefaultGroup)
	ticket_system.DELETE("delete_ticket_default_group", controllers.DeleteTicketDefaultGroup)
	ticket_system.PUT("update_ticket_default_group", controllers.UpdateTicketDefaultGroup)

	ticket_system.POST("create_ticket_default_timing", controllers.CreateTicketDefaultTiming)

	ticket_system.DELETE("delete_ticket_default_timing", controllers.DeleteTicketDefaultTiming)
	ticket_system.PUT("update_ticket_default_timing", controllers.UpdateTicketDefaultTiming)

	ticket_system.POST("create_ticket_default_type", controllers.CreateTicketDefaultType)
	ticket_system.GET("list_ticket_types", controllers.ListTicketType)
	ticket_system.GET("list_ticket_default_types", controllers.ListTicketDefaultType)
	ticket_system.DELETE("delete_ticket_default_type", controllers.DeleteTicketDefaultType)
	ticket_system.PUT("update_ticket_default_type", controllers.UpdateTicketDefaultType)

	ticket_system.POST("get_ticket_token", controllers.GetTicketToken)
	ticket_system.POST("create_token_ticket", controllers.CreateTokenTicket)
	ticket_system.POST("create_token_ticket_activity", controllers.CreateTokenTicketActivity)
	ticket_system.GET("list_token_ticket_details", controllers.ListTokenTicketDetail)

	ticket_system.PUT("update_token_ticket", controllers.UpdateTokenTicket)

	return r

}
