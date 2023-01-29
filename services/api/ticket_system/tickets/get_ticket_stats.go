package ticket_system

import (
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

const (
	YYYYMMDD = "2006-01-02"
)

func GetTicketStats(stats models.TicketStat) models.TicketStat {
	db := config.GetDB()

	var ticket_reviewer models.TicketReviewer
	var ticket_user models.TicketUser
	var ticket_id []uint
	t := time.Now()
	c := time.Now().Format(YYYYMMDD)
	y, _ := time.Parse(YYYYMMDD, c)

	if stats.AgentRmID != "7c6c1fe7-4a4d-4f3a-b432-b05ffdec3b44" {
		var ticket_users []uint
		db2 := config.GetCDB()
		var partner_user_rm []models.PartnerUserRmMapping
		var partner_user_rm_ids []string

		db2.Where("reporting_manager_id = ? and status = 'active'", stats.AgentRmID).Distinct("user_id").Find(&partner_user_rm).Pluck("user_id", &partner_user_rm_ids)

		db.Where("system_user_id IN ?", partner_user_rm_ids).Find(&ticket_user).Pluck("id", &ticket_users)

		db.Where("ticket_user_id IN ?", ticket_users).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	} else if stats.AgentID != "7c6c1fe7-4a4d-4f3a-b432-b05ffdec3b44" {
		db.Where("system_user_id = ?", stats.AgentID).First(&ticket_user)

		db.Where("ticket_user_id = ?", ticket_user.ID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
	} else {

		db.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
	}

	db = config.GetDB()

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "unresolved").Where("created_at BETWEEN ? and ?", y, y.AddDate(0, 10, 1)).Count(&stats.Unresolved)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("created_at BETWEEN ? and ?", y, y.AddDate(0, 10, 1)).Count(&stats.Closed)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "rejected").Where("created_at BETWEEN ? and ?", y, y.AddDate(0, 10, 1)).Count(&stats.Rejected)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ? AND tat BETWEEN ? AND ?", "unresolved", t.Format(YYYYMMDD), time.Now()).Where("created_at BETWEEN ? and ?", y, y.AddDate(0, 10, 1)).Count(&stats.DueToday)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "overdue").Where("created_at BETWEEN ? and ?", y, y.AddDate(0, 10, 1)).Count(&stats.Overdue)

	db.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "reassigned").Where("created_at BETWEEN ? and ?", y, y.AddDate(0, 10, 1)).Count(&stats.Reassigned)

	db.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "escalated").Where("created_at BETWEEN ? and ?", y, y.AddDate(0, 10, 1)).Count(&stats.Escalated)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("created_at BETWEEN ? and ?", y, y.AddDate(0, 10, 1)).Where("priority = 'high'").Count(&stats.HighPriority)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("expiry_date BETWEEN ? AND ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon)

	return stats
}
