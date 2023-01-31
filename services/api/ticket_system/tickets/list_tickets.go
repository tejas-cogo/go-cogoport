package ticket_system

import (
	"fmt"
	"strings"
	"time"
	"errors"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicket(filters models.TicketExtraFilter) ([]models.Ticket, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_user models.TicketUser
	var ticket_reviewer models.TicketReviewer
	var ticket_id []string

	const (
		YYYYMMDD = "2006-01-24"
	)

	var ticket []models.Ticket

	if filters.MyTicket != "" {
		db.Where("system_user_id = ?", filters.MyTicket).First(&ticket_user)
		db = db.Where("ticket_user_id = ?", ticket_user.ID)
	} else {
		if filters.AgentRmID != "" {
			var ticket_users []uint

			db2 := config.GetCDB()
			var partner_user_rm_mapping []models.PartnerUserRmMapping
			var partner_user_rm_ids []string

			db2.Where("reporting_manager_id = ? and status = ?", filters.AgentRmID, "active").Distinct("user_id").Find(&partner_user_rm_mapping).Pluck("user_id", &partner_user_rm_ids)

			db.Where("system_user_id IN ?", partner_user_rm_ids).Distinct("id").Find(&ticket_user).Pluck("id", &ticket_users)

			db.Where("ticket_user_id In ? or ticket_user_id = ? ", ticket_users, ticket_user.ID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

		} else if filters.AgentID != "" {
			db.Where("system_user_id = ?", filters.AgentID).First(&ticket_user)

			db.Where("ticket_user_id = ?", ticket_user.ID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
		} else {
			db.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
		}
		db = db.Where("id IN ?", ticket_id)
	}

	if filters.ID > 0 {
		tx = tx.Where("id = ?", filters.ID)
	}

	if filters.Type != "" {
		tx = tx.Where("type ilike ?", filters.Type)
	}

	if filters.QFilter != "" {

		tx = tx.Where("id::text ilike ? OR type ilike ?", filters.QFilter, "%"+filters.QFilter+"%")
	}

	if filters.Priority != "" {
		tx = tx.Where("priority = ?", filters.Priority)
	}

	if filters.IsExpiringSoon == "true" {
		x := time.Now()
		y := x.AddDate(0, 0, 10)
		fmt.Println(x, ", ", y)
		tx = tx.Where("expiry_date BETWEEN ? AND ?", x, y)
	}

	if filters.TicketUserID != 0 {
		tx = tx.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.TicketCreatedAt != "" {
		CreatedAt, _ := time.Parse(YYYYMMDD, filters.TicketCreatedAt)
		x := CreatedAt
		y := x.AddDate(0, 0, 1)
		tx = tx.Where("created_at BETWEEN ? AND ?", x, y)
	}

	if filters.ExpiryDate != "" {
		ExpiryDate, _ := time.Parse(YYYYMMDD, filters.ExpiryDate)
		x := ExpiryDate
		y := x.AddDate(0, 0, 1)
		tx = tx.Where("expiry_date BETWEEN ? AND ?", x, y)
	}

	if len(filters.Tags) != 0 {
		tx = tx.Where("tags && ?", "{"+strings.Join(filters.Tags, ",")+"}")
	}

	if filters.Status != "" {
		tx = tx.Where("status = ?", filters.Status)
	}

	db = db.Order("created_at desc").Order("expiry_date desc")

	tx = tx.Preload("TicketUser").Find(&ticket)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return ticket, tx, errors.New("Error Occurred!")
	}

	tx.Commit()
	return ticket, tx, err
}
