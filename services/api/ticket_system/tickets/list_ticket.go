package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicket(filters models.Ticket, sort models.Sort) ([]models.Ticket, *gorm.DB) {
	db := config.GetDB()

	var ticket []models.Ticket

	if filters.ID != 0 {
		db = db.Where("id = ?", filters.ID)
	}

	if filters.Type != "" {
		db = db.Where("type = ?", filters.Type)
	}

	if filters.Priority != "" {
		db = db.Where("priority = ?", filters.Priority)
	}

	if filters.Source != "" {
		db = db.Where("source = ?", filters.Source)
	}

	// if filters.Tags[0] != "" {
	// 	db = db.Where("? Like ANY(tags)", filters.Tags)
	// }

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	if sort.SortBy == "expiry_duration" && sort.SortType == "asc" {
		db = db.Order("expiry_duration asc").Order("created_at desc")
	} else if sort.SortBy == "expiry_duration" && sort.SortType == "desc" {
		db = db.Order("expiry_duration desc").Order("created_at desc")
	} else {
		db = db.Order("created_at desc").Order("expiry_duration desc")
	}

	db = db.Preload("TicketUser").Find(&ticket)

	return ticket, db
}
