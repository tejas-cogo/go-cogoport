package ticket_system

import (
	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketUser(filters models.TicketUser) ([]models.TicketUser, *gorm.DB) {
	db := config.GetDB()

	var ticket_user []models.TicketUser

	if filters.ID != 0 {
		db = db.Where("id = ?", filters.ID)
	}

	if filters.SystemUserID != uuid.Nil {
		db = db.Where("system_user_id = ?", filters.SystemUserID)
	}

	if filters.Name != "" {
		filters.Name = "%" + filters.Name + "%"
		db = db.Where("name LIKE ?", filters.Name)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Find(&ticket_user)

	return ticket_user, db
}
