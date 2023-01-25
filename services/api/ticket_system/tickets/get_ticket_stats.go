package ticket_system

import (
	"fmt"
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

	fmt.Println("stats.AgentRmID", stats)
	if stats.AgentRmID != "" {
		var ticket_users []uint
		var group_member models.GroupMember
		// db2 := config.GetCDB()
		// var partner_user_rm []models.PartnerUserRmMapping
		// var partner_user []models.PartnerUser

		// db2.Where("reporting_manager_id = ? and status = 'active'", stats.AgentRmID).Distinct("user_id").Find(&partner_user_rm)

		// db2.
		db.Where("system_user_id = ?", stats.AgentRmID).First(&ticket_user)

		db.Where("group_head_id = ?", ticket_user.ID).Distinct("ticket_user_id").Order("ticket_user_id").Find(&group_member).Pluck("ticket_user_id", &ticket_users)

		db.Where("ticket_user_id In ? or ticket_user_id = ? ", ticket_users, ticket_user.ID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	} else if stats.AgentID != "" {
		db.Where("system_user_id = ?", stats.AgentID).First(&ticket_user)

		db.Where("ticket_user_id = ?", ticket_user.ID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
	} else {

		db.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
	}

	fmt.Println("ticket", ticket_id, "ticket")

	db = config.GetDB()

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "unresolved").Count(&stats.Unresolved)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Count(&stats.Closed)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "rejected").Count(&stats.Rejected)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ? AND tat BETWEEN ? AND ?", "unresolved", t.Format(YYYYMMDD), time.Now()).Count(&stats.DueToday)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "overdue").Count(&stats.Overdue)

	db.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "reassigned").Count(&stats.Reassigned)

	db.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "escalated").Count(&stats.Escalated)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("priority = 'high'").Count(&stats.HighPriority)

	db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("expiry_date BETWEEN ? AND ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon)

	return stats
}
