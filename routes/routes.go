package routes

import (
	// "go-cogoport/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/tejas-cogo/go-cogoport/controllers/api"
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
	ticket_system.POST("create_ticket_user", controllers.CreateTicketUser)
	// ticket_system.DELETE("update_ticket_user_role", controllers.InactiveTicketUserRole)
	ticket_system.PUT("update_ticket_user", controllers.UpdateTicketUser)

	ticket_system.GET("list_ticket_details", controllers.ListTicketDetail)
	ticket_system.GET("list_tickets", controllers.ListTicket)
	ticket_system.GET("list_ticket_tags", controllers.ListTicketTag)
	ticket_system.GET("get_ticket_stats", controllers.GetTicketStats)
	ticket_system.GET("get_ticket_graph", controllers.GetTicketGraph)
	ticket_system.POST("create_ticket", controllers.CreateTicket)
	// ticket_system.DELETE("delete_ticket", controllers.DeleteTicket)
	ticket_system.PUT("update_ticket", controllers.UpdateTicket)

	// ticket_system.GET("list_ticket_activities", controllers.ListTicketActivity)
	ticket_system.POST("create_ticket_activity", controllers.CreateTicketActivity)
	// ticket_system.POST("delete_ticket_activity", controllers.DeleteTicketActivity)
	// ticket_system.POST("update_ticket_activity", controllers.UpdateTicketActivity)

	// ticket_system.GET("list_ticket_task", controllers.ListTicketTask)
	// ticket_system.POST("create_ticket_task", controllers.CreateTicketTask)
	// ticket_system.POST("delete_ticket_task", controllers.DeleteTicketTask)
	// ticket_system.POST("update_ticket_task", controllers.UpdateTicketTask)

	// ticket_system.GET("list_ticket_task_assignee", controllers.ListTicketTaskAssignee)
	// ticket_system.POST("create_ticket_task_assignee", controllers.CreateTicketTaskAssignee)
	// ticket_system.POST("delete_ticket_task_assignee", controllers.DeleteTicketTaskAssignee)
	// ticket_system.POST("update_ticket_task_assignee", controllers.UpdateTicketTaskAssignee)

	ticket_system.GET("list_ticket_spectators", controllers.ListTicketSpectator)
	ticket_system.POST("create_ticket_spectator", controllers.CreateTicketSpectator)
	ticket_system.DELETE("delete_ticket_spectator", controllers.DeleteTicketSpectator)
	// ticket_system.POST("update_ticket_spectator", controllers.UpdateTicketSpectator)

	// ticket_system.POST("create_ticket_audit", controllers.CreateTicketAudit)
	// ticket_system.POST("delete_ticket_audit", controllers.DeleteTicketAudit)
	// ticket_system.POST("update_ticket_audit", controllers.UpdateTicketAudit)

	// ticket_system.POST("create_ticket_reviewer", controllers.CreateTicketReviewer)
	ticket_system.POST("reassign_ticket_reviewer", controllers.ReassignTicketReviewer)
	// ticket_system.GET("list_ticket_reviewers", controllers.ListTicketReviewer)
	// ticket_system.POST("delete_ticket_reviewer", controllers.DeleteTicketReviewer)
	// ticket_system.POST("update_ticket_reviewer", controllers.UpdateTicketReviewer)

	ticket_system.POST("create_ticket_default_group", controllers.CreateTicketDefaultGroup)
	ticket_system.DELETE("delete_ticket_default_group", controllers.DeleteTicketDefaultGroup)
	ticket_system.PUT("update_ticket_default_group", controllers.UpdateTicketDefaultGroup)

	ticket_system.POST("create_ticket_default_timing", controllers.CreateTicketDefaultTiming)
	// ticket_system.GET("list_ticket_default_timing", controllers.ListTicketDefaultTiming)
	ticket_system.DELETE("delete_ticket_default_timing", controllers.DeleteTicketDefaultTiming)
	ticket_system.PUT("update_ticket_default_timing", controllers.UpdateTicketDefaultTiming)

	ticket_system.POST("create_ticket_default_type", controllers.CreateTicketDefaultType)
	ticket_system.GET("list_ticket_types", controllers.ListTicketType)
	ticket_system.GET("list_ticket_default_types", controllers.ListTicketDefaultType)
	ticket_system.DELETE("delete_ticket_default_type", controllers.DeleteTicketDefaultType)
	ticket_system.PUT("update_ticket_default_type", controllers.UpdateTicketDefaultType)

	ticket_system.POST("create_ticket_token", controllers.CreateTicketToken)
	ticket_system.POST("create_token_ticket", controllers.CreateTokenTicket)
	ticket_system.POST("create_token_ticket_activity", controllers.CreateTokenTicketActivity)
	ticket_system.GET("list_token_ticket_details", controllers.ListTokenTicketDetail)
	// ticket_system.DELETE("delete_ticket_token", controllers.DeleteTicketToken)
	ticket_system.PUT("update_token_ticket", controllers.UpdateTokenTicket)

	return r

}
