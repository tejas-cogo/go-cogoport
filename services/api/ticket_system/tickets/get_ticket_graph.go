package ticket_system

import (
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetTicketGraph(graph models.TicketGraph) models.TicketGraph {
	db := config.GetDB()

	var ticket_reviewer models.TicketReviewer
	var ticket_user models.TicketUser
	var ticket_id []uint

	const (
		DateTime = "2006-01-02"
	)

	if graph.AgentRmID != "7c6c1fe7-4a4d-4f3a-b432-b05ffdec3b44" {
		var ticket_users []uint
		db2 := config.GetCDB()
		var partner_user_rm []models.PartnerUserRmMapping
		var partner_user_rm_ids []string

		db2.Where("reporting_manager_id = ? and status = 'active'", graph.AgentRmID).Distinct("user_id").Find(&partner_user_rm).Pluck("user_id", &partner_user_rm_ids)

		db.Where("system_user_id IN ?", partner_user_rm_ids).Find(&ticket_user).Pluck("id", &ticket_users)

		db.Where("ticket_user_id IN ?", ticket_users).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	} else if graph.AgentID != "7c6c1fe7-4a4d-4f3a-b432-b05ffdec3b44" {
		db.Where("system_user_id = ?", graph.AgentID).First(&ticket_user)

		db.Where("ticket_user_id = ?", ticket_user.ID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
	} else {

		db.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)
	}

	db = config.GetDB()

	var x time.Time
	var y time.Time
	t := time.Now().Format(DateTime)
	y, _ = time.Parse(DateTime, t)

	for i := 1; i <= 6; i++ {

		var stats models.TicketStat
		x = y
		y = x.Add(time.Hour * 4)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("updated_at BETWEEN ?  AND ?", x, y).Count(&stats.Closed)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "unresolved").Where("created_at BETWEEN ?  AND ?", x, y).Count(&stats.Open)

		switch x.Weekday() {
		case 1:
			graph.TodayClosed.First = stats.Closed
			graph.TodayOpen.First = stats.Open
		case 2:
			graph.TodayClosed.Second = stats.Closed
			graph.TodayOpen.Second = stats.Open
		case 3:
			graph.TodayClosed.Third = stats.Closed
			graph.TodayOpen.Third = stats.Open
		case 4:
			graph.TodayClosed.Fourth = stats.Closed
			graph.TodayOpen.Fourth = stats.Open
		case 5:
			graph.TodayClosed.Fifth = stats.Closed
			graph.TodayOpen.Fifth = stats.Open
		case 6:
			graph.TodayClosed.Sixth = stats.Closed
			graph.TodayOpen.Sixth = stats.Open
		}

	}

	weekday := time.Now().Weekday()

	y, _ = time.Parse(DateTime, t)

	t1 := int(weekday)

	t1 = -t1 + 1

	y = x.AddDate(0, 0, t1)

	for i := 1; i <= 7; i++ {

		var stats models.TicketStat
		x = y
		y = x.AddDate(0, 0, 1)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("updated_at BETWEEN ?  AND ?", x, y).Count(&stats.Closed)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "unresolved").Where("created_at BETWEEN ?  AND ?", x, y).Count(&stats.Open)

		switch x.Weekday() {
		case 1:
			graph.WeekClosed.Monday = stats.Closed
			graph.WeekOpen.Monday = stats.Open
		case 2:
			graph.WeekClosed.Tuesday = stats.Closed
			graph.WeekOpen.Tuesday = stats.Open
		case 3:
			graph.WeekClosed.Wednesday = stats.Closed
			graph.WeekOpen.Wednesday = stats.Open
		case 4:
			graph.WeekClosed.Thursday = stats.Closed
			graph.WeekOpen.Thursday = stats.Open
		case 5:
			graph.WeekClosed.Friday = stats.Closed
			graph.WeekOpen.Friday = stats.Open
		case 6:
			graph.WeekClosed.Saturday = stats.Closed
			graph.WeekOpen.Saturday = stats.Open
		case 7:
			graph.WeekClosed.Sunday = stats.Closed
			graph.WeekOpen.Sunday = stats.Open
		}

	}

	return graph
}
