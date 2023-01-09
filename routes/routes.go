package routes

import (
	// "go-cogoport/middlewares"
	"github.com/gin-gonic/gin"
	controllers "github.com/tejas-cogo/go-cogoport/controllers/api"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	// user := r.Group("/api/user")

	// user.GET("user_list", controllers.UserList)
	// user.POST("create_user", controllers.CreateUser)
	// user.POST("delete_user", controllers.DeleteUser)
	// user.POST("update_user", controllers.UpdateUser)

	ticket_system := r.Group("/api/tickets")

	// ticket_system.GET("list_group", controllers.ListGroup)
	ticket_system.POST("create_group", controllers.CreateGroup)
	ticket_system.POST("delete_group", controllers.DeleteGroup)
	ticket_system.POST("update_group", controllers.UpdateGroup)

	ticket_system.POST("create_group_member", controllers.CreateGroupMember)
	ticket_system.POST("delete_group_member", controllers.DeleteGroupMember)
	ticket_system.POST("update_group_member", controllers.UpdateGroupMember)

	// ticket_system.GET("list_role", controllers.ListRole)
	ticket_system.POST("create_role", controllers.CreateRole)
	ticket_system.POST("delete_role", controllers.DeleteRole)
	ticket_system.POST("update_role", controllers.UpdateRole)

	ticket_system.POST("create_ticket_user", controllers.CreateTicketUser)
	ticket_system.POST("delete_ticket_user", controllers.DeleteTicketUser)
	ticket_system.POST("update_ticket_user", controllers.UpdateTicketUser)

	// ticket_system.GET("list_ticket", controllers.ListTicket)
	ticket_system.POST("create_ticket", controllers.CreateTicket)
	ticket_system.POST("delete_ticket", controllers.DeleteTicket)
	ticket_system.POST("update_ticket", controllers.UpdateTicket)

	// ticket_system.GET("list_ticket_activity", controllers.ListTicketActivity)
	ticket_system.POST("create_ticket_activity", controllers.CreateTicketActivity)
	ticket_system.POST("delete_ticket_activity", controllers.DeleteTicketActivity)
	ticket_system.POST("update_ticket_activity", controllers.UpdateTicketActivity)

	ticket_system.GET("list_ticket_task", controllers.ListTicketTask)
	ticket_system.POST("create_ticket_task", controllers.CreateTicketTask)
	ticket_system.POST("delete_ticket_task", controllers.DeleteTicketTask)
	ticket_system.POST("update_ticket_task", controllers.UpdateTicketTask)

	ticket_system.GET("list_ticket_task_assignee", controllers.ListTicketTaskAssignee)
	ticket_system.POST("create_ticket_task_assignee", controllers.CreateTicketTaskAssignee)
	ticket_system.POST("delete_ticket_task_assignee", controllers.DeleteTicketTaskAssignee)
	ticket_system.POST("update_ticket_task_assignee", controllers.UpdateTicketTaskAssignee)

	ticket_system.GET("list_ticket_spectator", controllers.ListTicketSpectator)
	ticket_system.POST("create_ticket_spectator", controllers.CreateTicketSpectator)
	ticket_system.POST("delete_ticket_spectator", controllers.DeleteTicketSpectator)
	ticket_system.POST("update_ticket_spectator", controllers.UpdateTicketSpectator)

	// ticket_system.GET("list_ticket_audit", controllers.ListTicketAudit)
	ticket_system.POST("create_ticket_audit", controllers.CreateTicketAudit)
	ticket_system.POST("delete_ticket_audit", controllers.DeleteTicketAudit)
	ticket_system.POST("update_ticket_audit", controllers.UpdateTicketAudit)

	// ticket_system.GET("list_ticket_reviewer", controllers.CreateTicketReviewer)
	ticket_system.POST("delete_ticket_reviewer", controllers.DeleteTicketReviewer)
	ticket_system.POST("update_ticket_reviewer", controllers.UpdateTicketReviewer)

	ticket_system.POST("create_ticket_default_group", controllers.CreateTicketDefaultGroup)
	ticket_system.POST("delete_ticket_default_group", controllers.DeleteTicketDefaultGroup)
	ticket_system.POST("update_ticket_default_group", controllers.UpdateTicketDefaultGroup)

	ticket_system.POST("create_ticket_default_timing", controllers.CreateTicketDefaultTiming)
	//ticket_system.GET("list_ticket_default_timing", controllers.ListTicketDefaultTiming)
	ticket_system.POST("delete_ticket_default_timing", controllers.DeleteTicketDefaultTiming)
	ticket_system.POST("update_ticket_default_timing", controllers.UpdateTicketDefaultTiming)

	ticket_system.POST("create_ticket_default_type", controllers.CreateTicketDefaultType)
	ticket_system.POST("delete_ticket_default_type", controllers.DeleteTicketDefaultType)
	ticket_system.POST("update_ticket_default_type", controllers.UpdateTicketDefaultType)

	return r

}
