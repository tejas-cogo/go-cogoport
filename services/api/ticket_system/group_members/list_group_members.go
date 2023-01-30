package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListGroupMember(filters models.FilterGroupMember) ([]models.GroupMember, *gorm.DB) {
	db := config.GetDB()

	var group_members []models.GroupMember

	if filters.GroupID > 0 {
		db = db.Where("group_id = ?", filters.GroupID)
	}

	if filters.TicketUserID > 0 {
		db = db.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	if filters.NotPresentTicketUserID > 0 {
		db = db.Where("ticket_user_id != ?", filters.NotPresentTicketUserID)
	}

	db = db.Order("hierarchy_level desc").Order("active_ticket_count asc").Preload("TicketUser").Preload("Group").Find(&group_members)

	return group_members, db
}
