package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListGroupMember(filters models.GroupMember) ([]models.GroupMember, *gorm.DB) {
	db := config.GetDB()

	var group_members []models.GroupMember

	if filters.GroupID != 0 {
		db = db.Where("group_id = ?", filters.GroupID)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	} else {
		db = db.Where("status = ?", "active")
	}

	db.Order("hierarchy_level desc").Order("active_ticket_count asc")

	db.Find(&group_members)

	return group_members, db
}
