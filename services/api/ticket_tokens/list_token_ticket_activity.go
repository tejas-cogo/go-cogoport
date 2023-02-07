package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTokenTicketActivity(token_filter models.TokenFilter) ([]models.TicketActivity, *gorm.DB, error) {
	db := config.GetDB()

	var ticket_token models.TicketToken
	var ticket_activity []models.TicketActivity

	var err error

	if err = db.Where("ticket_token = ? and status= ?", token_filter.TicketToken, "utilized").First(&ticket_token).Error; err != nil {
		return ticket_activity, db, errors.New("token not found!")
	}

	if ticket_token.TicketID > 0 {
		db = db.Where("ticket_id = ?", ticket_token.TicketID)
	}

	db = db.Order("created_at desc").Preload("Ticket").Find(&ticket_activity)

	return ticket_activity, db, err
}
