package ticket_system

import (
	"fmt"
	"strings"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicket(filters models.TicketExtraFilter) ([]models.Ticket, *gorm.DB) {
	db := config.GetDB()
	var ticket_user models.TicketUser
	var ticket_reviewer models.TicketReviewer
	var ticket_id []string

	const (
		YYYYMMDD = "2006-01-24"
	)

	var ticket []models.Ticket

	// if filters.AgentRmID != "" {
	// 	var ticket_users []uint
	// 	var group_member models.GroupMember

	// 	db.Where("system_user_id = ?", filters.AgentRmID).First(&ticket_user)

	// 	db.Where("group_head_id = ?", ticket_user.ID).Distinct("ticket_user_id").Order("ticket_user_id").Find(&group_member).Pluck("ticket_user_id", &ticket_users)

	// 	db.Where("ticket_user_id In ? or ticket_user_id = ? ", ticket_users, ticket_user.ID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	// } else
	if filters.AgentID != "" {
		db.Where("system_user_id = ?", filters.AgentID).First(&ticket_user)

		db.Where("ticket_user_id = ?", ticket_user.ID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
	} else {

		db.Distinct("id").Order("id").Find(&ticket).Pluck("id", &ticket_id)
	}

	if filters.ID > 0 {
		db = db.Where("id = ?", filters.ID)
	}

	if filters.Type != "" {
		db = db.Where("type = ?", filters.Type)
	}

	if filters.Priority != "" {
		db = db.Where("priority = ?", filters.Priority)
	}

	if filters.IsExpiringSoon == "true" {
		x := time.Now()
		y := x.AddDate(0, 0, 10)
		fmt.Println(x, ", ", y)
		db = db.Where("expiry_date BETWEEN ? AND ?", x, y)
	}

	if filters.TicketUserID != 0 {
		db = db.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.TicketUserID != 0 {
		db = db.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.TicketCreatedAt != "" {
		CreatedAt, _ := time.Parse(YYYYMMDD, filters.TicketCreatedAt)
		x := CreatedAt
		y := x.AddDate(0, 0, 1)
		db = db.Where("created_at BETWEEN ? AND ?", x, y)
	}

	if filters.ExpiryDate != "" {
		ExpiryDate, _ := time.Parse(YYYYMMDD, filters.ExpiryDate)
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

	db = db.Where("id IN ?", ticket_id)

	// if sort.SortBy == "expiry_duration" && sort.SortType == "asc" {
	// 	db = db.Order("expiry_date asc").Order("created_at desc")
	// } else if sort.SortBy == "expiry_duration" && sort.SortType == "desc" {
	// 	db = db.Order("expiry_date desc").Order("created_at desc")
	// } else {
	db = db.Order("created_at desc").Order("expiry_date desc")

	db = db.Preload("TicketUser").Find(&ticket)

	return ticket, db
}
