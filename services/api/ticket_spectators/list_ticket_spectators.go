package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketSpectator(filters models.TicketSpectator) ([]models.TicketSpectator, *gorm.DB, error) {
	db := config.GetDB()

	var err error

	var ticket_spectator []models.TicketSpectator

	if filters.TicketID != 0 {
		db = db.Where("ticket_id = ?", filters.TicketID)
	}

	if filters.UserID != "" {
		db = db.Where("ticket_user_id = ?", filters.UserID)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Preload("TicketUser").Find(&ticket_spectator)



	return ticket_spectator, db, err
}
