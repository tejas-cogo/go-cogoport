package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
	"errors"
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
		tx = tx.Where("status = ?", filters.Status)
	}

	if filters.NotPresentTicketUserID > 0 {
		tx = tx.Where("ticket_user_id != ?", filters.NotPresentTicketUserID)
	}

	tx = tx.Order("hierarchy_level desc").Order("active_ticket_count asc")

	if filters.GroupMemberName != "" {
		tx = tx.Preload("TicketUser", "name = ?", filters.GroupMemberName)
	} else {
		tx = tx.Preload("TicketUser")
	}

	tx = tx.Preload("Group").Find(&group_members)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return group_members, tx, errors.New("Error Occurred!")
	}

	tx.Commit()
	return group_members, tx, err
}
