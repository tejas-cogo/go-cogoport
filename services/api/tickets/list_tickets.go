package api

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/constants"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicket(filters models.TicketExtraFilter) ([]models.Ticket, *gorm.DB) {
	db := config.GetDB()

	var ticket_reviewer models.TicketReviewer
	var ticket_id []string

	var ticket []models.Ticket

	if filters.MyTicket != uuid.Nil {
		db = db.Where("user_id = ?", filters.MyTicket).Distinct("id").Order("id").Find(&ticket).Pluck("id", &ticket_id)

	} else {
		if filters.AgentRmID != uuid.Nil {

			db.Where("manager_rm_ids && '(?)' or user_id = ?", filters.AgentRmID, filters.AgentRmID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

		} else if filters.AgentID != uuid.Nil {

			db.Where("user_id = ?", filters.AgentID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

		} else {

			db.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

		}
	}

	db = db.Where("id IN ?", ticket_id)

	if filters.ID > 0 {
		db = db.Where("id = ?", filters.ID)
	}

	if filters.UserID != uuid.Nil {
		db = db.Where("id = ?", filters.UserID)
	}

	if filters.Type != "" {
		db = db.Where("type ilike ?", filters.Type)
	}

	if filters.QFilter != "" {

		db = db.Where("id::text ilike ? OR type ilike ?", filters.QFilter, "%"+filters.QFilter+"%")
	}

	if filters.Priority != "" {
		db = db.Where("priority = ?", filters.Priority)
	}

	if filters.ExpiringSoon == "true" {
		x := time.Now()
		y := x.AddDate(0, 0, 1)
		db = db.Where("expiry_date BETWEEN ? AND ?", x, y)
	}

	if filters.TicketCreatedAt != "" {
		CreatedAt, _ := time.Parse(constants.DateTimeFormat(), filters.TicketCreatedAt)
		x := CreatedAt
		y := x.AddDate(0, 0, 1)
		db = db.Where("created_at BETWEEN ? AND ?", x, y)
	}

	if filters.ExpiryDate != "" {
		ExpiryDate, _ := time.Parse(constants.DateTimeFormat(), filters.ExpiryDate)
		x := ExpiryDate
		y := x.AddDate(0, 0, 1)
		db = db.Where("expiry_date BETWEEN ? AND ?", x, y)
	}

	if len(filters.Tags) != 0 {
		db = db.Where("tags && ?", "{"+strings.Join(filters.Tags, ",")+"}")
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Order("created_at desc").Order("expiry_date desc")

	db = db.Preload("TicketUser").Find(&ticket)

	return ticket, db
}
