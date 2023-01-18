package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetTicketStats(stats models.TicketStat) uint {
	db := config.GetDB()

	var ticket_reviewer models.TicketReviewer
	// var ticket_activity models.TicketActivity
	var ticket models.Ticket

	if stats.PerformedByID != 0 {
		db = db.Where("ticket_user_id = ?", stats.PerformedByID)
		var ticket_id []uint
		var status []string
		db.Distinct("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
		db.Where("id IN ?", ticket_id).Find(&ticket).Pluck("status", &status)
	}

	return stats.PerformedByID
}
