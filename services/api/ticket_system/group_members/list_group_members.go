package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListGroupMember(filters models.FilterGroupMember) ([]models.GroupMember, *gorm.DB) {
	db := config.GetDB()

	var group_members []models.GroupMember
	var ticket_user []models.TicketUser

	if filters.AgentRmID != "" || filters.AgentID != "" {
		var ticket_users []uint
		if filters.AgentRmID != "" {

			db2 := config.GetCDB()
			var partner_user_rm_mapping []models.PartnerUserRmMapping
			var partner_user_rm_ids []string

			db2.Where("reporting_manager_id = ? and status = ?", filters.AgentRmID, "active").Distinct("user_id").Find(&partner_user_rm_mapping).Pluck("user_id", &partner_user_rm_ids)

			db.Where("system_user_id IN ?", partner_user_rm_ids).Distinct("id").Find(&ticket_user).Pluck("id", &ticket_users)

		} else if filters.AgentID != "" {
			db.Where("system_user_id = ?", filters.AgentID).Find(&ticket_user).Pluck("id", &ticket_users)
		}
		db = db.Where("ticket_user_id IN ?", ticket_users)
	}

	if filters.GroupID > 0 {
		db = db.Where("group_id = ?", filters.GroupID)
	}

	if filters.TicketUserID > 0 {
		db = db.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.Status != "" {

		db = db.Where("group_members.status = ?", filters.Status)

	}

	if filters.NotPresentTicketUserID > 0 {
		db = db.Where("ticket_user_id != ?", filters.NotPresentTicketUserID)
	}

	db = db.Order("hierarchy_level desc").Order("active_ticket_count asc")

	if filters.GroupMemberName != "" {
		filters.GroupMemberName = "%" + filters.GroupMemberName + "%"

		db = db.Joins("Inner Join ticket_users on ticket_users.id = group_members.ticket_user_id and ticket_users.name iLike ?", filters.GroupMemberName)
	}
	db = db.Preload("TicketUser")

	db = db.Preload("Group").Find(&group_members)

	return group_members, db
}
