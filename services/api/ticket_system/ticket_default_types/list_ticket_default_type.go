package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketDefaultType(filters models.TicketDefaultType) ([]models.TicketDefaultType, *gorm.DB) {
	db := config.GetDB()

	var ticket_default_type []models.TicketDefaultType

	if filters.TicketType != "" {
		db = db.Where("ticket_type LIKE ?", filters.TicketType)
	}

	db.Find(&ticket_default_type)

	return ticket_default_type, db
}
