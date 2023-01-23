package ticket_system

import (
	"fmt"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicket(filters models.Ticket, sort models.Sort, tags string) ([]models.Ticket, *gorm.DB) {
	db := config.GetDB()

	const (
		YYYYMMDD = "2006-01-24"
	)

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

	if filters.TicketUserID != 0 {
		db = db.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.TicketUserID != 0 {
		db = db.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if !time.Time.IsZero(filters.CreatedAt) {

		x := filters.CreatedAt
		y := x.AddDate(0, 0, 1)
		fmt.Println(x, ", ", y)
		db = db.Where("created_at BETWEEN ? AND ?", x, y)
	}

	if !time.Time.IsZero(filters.ExpiryDate) {
		x := filters.ExpiryDate
		y := x.AddDate(0, 0, 1)
		db = db.Where("expiry_date BETWEEN ? AND ?", x, y)
	}

	if tags != "" {
		db = db.Where("? Like ANY(tags)", tags)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	if sort.SortBy == "expiry_duration" && sort.SortType == "asc" {
		db = db.Order("expiry_date asc").Order("created_at desc")
	} else if sort.SortBy == "expiry_duration" && sort.SortType == "desc" {
		db = db.Order("expiry_date desc").Order("created_at desc")
	} else {
		db = db.Order("created_at desc").Order("expiry_date desc")
	}

	db = db.Preload("TicketUser").Find(&ticket)

	return ticket, db
}
