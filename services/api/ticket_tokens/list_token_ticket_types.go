package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTokenTicketType(body models.TokenFilter) ([]models.TicketDefaultType, *gorm.DB) {

	db := config.GetDB()

	var ticket_token models.TicketToken

	db.Where("ticket_token = ? and status = ?", body.TicketToken, "utilized").First(&ticket_token)

	var ticket_default_type []models.TicketDefaultType

	if body.TicketType != "" {
		body.TicketType = "%" + body.TicketType + "%"
		db = db.Where("ticket_type iLike ?", body.TicketType)
	}
	if body.Status != "" {
		db = db.Where("status = ?", body.Status)
	}
	db = db.Where("id != ?", 1)
	db = db.Order("created_at desc").Find(&ticket_default_type)

	return ticket_default_type, db
}
