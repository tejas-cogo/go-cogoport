package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
	"errors"
)

func ListTicketSpectator(filters models.TicketSpectator) ([]models.TicketSpectator, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_spectator []models.TicketSpectator

	if filters.TicketID != 0 {
		tx = tx.Where("ticket_id = ?", filters.TicketID)
	}

	if filters.TicketUserID != 0 {
		tx = tx.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.Status != "" {
		tx = tx.Where("status = ?", filters.Status)
	}

	tx = tx.Preload("TicketUser").Find(&ticket_spectator)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return ticket_spectator, tx, errors.New("Error Occurred!")
	}

	tx.Commit()

	return ticket_spectator, tx, err
}
