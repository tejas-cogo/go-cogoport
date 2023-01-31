package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListGroupMember(filters models.FilterGroupMember) ([]models.GroupMember, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var group_members []models.GroupMember

	if filters.GroupID > 0 {
		tx = tx.Where("group_id = ?", filters.GroupID)
	}

	if filters.TicketUserID > 0 {
		tx = tx.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.Status != "" {

		tx = tx.Where("group_members.status = ?", filters.Status)

	}

	if filters.NotPresentTicketUserID > 0 {
		tx = tx.Where("ticket_user_id != ?", filters.NotPresentTicketUserID)
	}

	tx = tx.Order("hierarchy_level desc").Order("active_ticket_count asc")

	if filters.GroupMemberName != "" {

		tx = tx.Joins("Inner Join ticket_users on ticket_users.id = group_members.ticket_user_id and ticket_users.name iLike ?", filters.GroupMemberName)
	}
	tx = tx.Preload("TicketUser")

	tx = tx.Preload("Group").Find(&group_members)

	tx.Commit()
	return group_members, tx, err
}
