package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketType(filters models.TicketDefaultType) ([]models.TicketDefaultType, *gorm.DB) {

	db := config.GetDB()

	var ticket_default_type []models.TicketDefaultType

	if filters.TicketType != "" {
		filters.TicketType = "%" + filters.TicketType + "%"
		db = db.Where("ticket_type iLike ?", filters.TicketType)
	}
	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}
	db = db.Where("id != ?", 1)
	db = db.Order("created_at desc").Find(&ticket_default_type)

	return ticket_default_type, db
}
