package api

import (
	"strings"
	"time"

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

	if filters.MyTicket != "" {
		db.Where("user_id = ?", filters.MyTicket).Distinct("id").Order("id").Find(&ticket).Pluck("id", &ticket_id)

	} else {
		if filters.AgentRmID != "" {

			db.Where("manager_rm_ids && '(?)' or user_id = ? and status = ?", filters.AgentRmID, filters.AgentRmID, "active").Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

		} else if filters.AgentID != "" {

			db.Where("user_id =  ? and status = ?", filters.AgentID, "active").Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

		} else {

			db.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

		}

	}

	if filters.Closure == true {
		db2 := config.GetDB()
		db2.Model(&models.Ticket{}).Distinct("tickets.id").Joins("inner join ticket_default_types on ticket_default_types.id = tickets.ticket_default_type_id and ticket_default_types.status = ?",
			"active").Where("tickets.status = ? and ticket_default_types.closure_authorizer &&  ?  and tickets.id IN ?", "pending", "{"+filters.ClosureID+"}", ticket_id).Pluck("tickets.id", &ticket_id)

	}

	db = db.Where("id IN ?", ticket_id)

	if filters.ID > 0 {
		db = db.Where("id = ?", filters.ID)
	}

	if filters.UserID != "" {
		db = db.Where("id = ?", filters.UserID)
	}

	if filters.Type != "" {
		db = db.Where("type ilike ?", filters.Type)
	}

	if filters.Category != "" {
		db = db.Where("category ilike ?", filters.Category)
	}

	if filters.QFilter != "" {

		db = db.Where("id::text ilike ? OR type ilike ? OR category ilike ?", "%"+filters.QFilter+"%", "%"+filters.QFilter+"%", "%"+filters.QFilter+"%")
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
	} else if filters.Statuses != "" {

		data := strings.Split(filters.Statuses, ",")

		db = db.Where("status IN (?)", data)
	}

	db = db.Order("created_at desc").Order("expiry_date desc").Find(&ticket)

	return ticket, db
}
