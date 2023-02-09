package api

import (
	"strings"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/constants"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetTicketStats(stats models.TicketStat) (models.TicketStat, error) {
	db := config.GetDB()
	var err error

	var ticket_reviewer []models.TicketReviewer
	var ticket_id []uint
	t := time.Now()

	if stats.AgentRmID != "" {

		db.Where("manager_rm_ids && '(?)' or user_id = ?", stats.AgentRmID, stats.AgentRmID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	} else if stats.AgentID != "" {

		db.Where("user_id = ?", stats.AgentID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	} else {

		db.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	}

	if stats.StartDate != "" {
		start_date, _ := time.Parse(constants.DateTimeFormat(), stats.StartDate)
		end_date, _ := time.Parse(constants.DateTimeFormat(), stats.EndDate)
		end_date = end_date.AddDate(0, 0, 1)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "unresolved").Where("created_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Unresolved)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("updated_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Closed)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "rejected").Where("updated_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Rejected)

		db.Model(&models.Ticket{}).Where("id IN ? and created_at::date = ? and status = ?", ticket_id, time.Now().Format(constants.DateTimeFormat()), "unresolved").Count(&stats.DueToday)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "overdue").Where("updated_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Overdue)

		db.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "reassigned").Where("created_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Reassigned)

		db.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "escalated").Where("created_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Escalated)

	} else {
		var ticket models.Ticket
		if stats.ExpiryDate != "" || stats.TicketCreatedAt != "" || len(stats.Tags) != 0 || stats.UserID != "" || stats.QFilter != "" {
			if stats.ExpiryDate != "" {
				ExpiryDate, _ := time.Parse(constants.DateTimeFormat(), stats.ExpiryDate)
				x := ExpiryDate
				y := x.AddDate(0, 0, 1)
				db = db.Where("expiry_date BETWEEN ? AND ?", x, y)
			}

			if stats.TicketCreatedAt != "" {
				CreatedAt, _ := time.Parse(constants.DateTimeFormat(), stats.TicketCreatedAt)
				x := CreatedAt
				y := x.AddDate(0, 0, 1)
				db = db.Where("created_at BETWEEN ? AND ?", x, y)
			}

			if len(stats.Tags) != 0 {
				db = db.Where("tags && ?", "{"+strings.Join(stats.Tags, ",")+"}")
			}

			if stats.UserID != "" {
				db = db.Where("user_id = ?", stats.UserID)
			}
			if stats.Category != "" {
				db = db.Where("category ilike ?", stats.Category)
			}

			if stats.QFilter != "" {
				db = db.Where("id::text ilike ? OR type ilike ? OR category ilike ?", "%"+stats.QFilter+"%", "%"+stats.QFilter+"%", "%"+stats.QFilter+"%")
			}
			db.Where("id IN ?", ticket_id).Distinct("id").Find(&ticket).Pluck("id ", &ticket_id)

		}

		db = config.GetDB()

		db.Where("id IN ? and status = ?", ticket_id, "unresolved").Model(&models.Ticket{}).Count(&stats.Unresolved)

		db.Model(&models.Ticket{}).Where("id In ? and priority = ? and status = ?", ticket_id, "high", "unresolved").Count(&stats.HighPriority)

		db.Model(&models.Ticket{}).Where("id In ? and created_at::date = ? and status = ?", ticket_id, time.Now().Format(constants.DateTimeFormat()), "unresolved").Count(&stats.DueToday)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Count(&stats.Closed)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "rejected").Count(&stats.Rejected)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "overdue").Count(&stats.Overdue)

		db.Model(&models.TicketActivity{}).Where("id IN ?", ticket_id).Where("status = ?", "reassigned").Count(&stats.Reassigned)

		db.Model(&models.TicketActivity{}).Where("id IN ?", ticket_id).Where("status = ?", "escalated").Count(&stats.Escalated)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("expiry_date BETWEEN ? AND ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon)
	}

	db.Commit()
	return stats, err
}
