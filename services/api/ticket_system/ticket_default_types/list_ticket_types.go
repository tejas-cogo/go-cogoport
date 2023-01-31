package ticket_system

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketType(filters models.TicketDefaultType) ([]models.TicketDefaultType, *gorm.DB, error) {

	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_type []models.TicketDefaultType

	if filters.TicketType != "" {
		filters.TicketType = "%" + filters.TicketType + "%"
		tx = tx.Where("ticket_type iLike ?", filters.TicketType)
	}
	if filters.Status != "" {
		tx = tx.Where("status = ?", filters.Status)
	}
	tx = tx.Where("id != ?", 1)
	tx = tx.Order("created_at desc").Find(&ticket_default_type)

	if err := tx.Error; err != nil {
		tx.Rollback()
		return ticket_default_type, tx, errors.New("Error Occurred!")
	}

	tx.Commit()
	return ticket_default_type, tx, err
}
