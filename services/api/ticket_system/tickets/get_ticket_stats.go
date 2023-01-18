package ticket_system

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetTicketStats(stats models.TicketStat) models.TicketStat {
	db := config.GetDB()

	var ticket_reviewer models.TicketReviewer
	var ticket_user models.TicketUser
	var ticket_id []uint
	var ticket_user_id []uint
	t := time.Now()
	const (
		YYYYMMDD = "2006-01-02"
	)

	if stats.PerformedByID != uuid.Nil {
		db = db.Where("system_user_id = ?", stats.PerformedByID).First(&ticket_user).Pluck("ticket_user_id", &ticket_user_id)
	}

	if stats.PerformedByID != uuid.Nil {
		db = db.Where("ticket_user_id = ?", ticket_user_id)
		
		db.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
		// db = db.Where("ticket_id IN ?", &ticket_id)
		// db.Where("status = ?", "reassigned").Model(&ticket_activity).Count(&stats.Reassigned)
		// db.Where("status = ?", "escalated").Model(&ticket_activity).Count(&stats.Escalated)
		db := config.GetDB()
		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "open").Count(&stats.Open)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Count(&stats.Closed)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "rejected").Count(&stats.Rejected)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ? AND tat BETWEEN ? AND ?", "open", t.Format(YYYYMMDD), time.Now()).Count(&stats.DueToday)

		fmt.Println("FVDC", t.Format(YYYYMMDD), "DCS")
		fmt.Println("FVDC", t, "DCS")

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "overdue").Count(&stats.Overdue)

		db.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "reassigned").Count(&stats.Reassigned)

		db.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "escalated").Count(&stats.Escalated)
	}

	return stats
}
