package ticket_system

import (
	"fmt"
	"time"
	"errors"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetTicketGraph(graph models.TicketGraph) (models.TicketGraph,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_reviewer []models.TicketReviewer
	var ticket_user models.TicketUser
	var ticket_id []uint
	var ticket_users []uint

	const (
		DateTime = "2006-01-02"
	)

	if graph.AgentRmID != "" {

		db2 := config.GetCDB()
		tx2 := db2.Begin()

		var partner_user_rm_mapping []models.PartnerUserRmMapping
		var partner_user_rm_ids []string

		if err := tx2.Where("reporting_manager_id = ? and status = ?", graph.AgentRmID, "active").Distinct("user_id").Find(&partner_user_rm_mapping).Pluck("user_id", &partner_user_rm_ids).Error; err != nil {
			tx.Rollback()
			return graph, errors.New("Error Occurred!")
		}

		fmt.Println("partner_user_rm_ids", partner_user_rm_ids)

		if err := tx.Where("system_user_id IN ?", partner_user_rm_ids).Distinct("id").Find(&ticket_user).Pluck("id", &ticket_users).Error; err != nil {
			tx.Rollback()
			return graph, errors.New("Error Occurred!")
		}

		if err := tx.Where("ticket_user_id IN ?", ticket_users).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id).Error; err != nil {
			tx.Rollback()
			return graph, errors.New("Error Occurred!")
		}

	} else if graph.AgentID != "" {
		if err := tx.Where("system_user_id = ?", graph.AgentID).Distinct("id").Find(&ticket_user).Pluck("id", &ticket_users).Error; err != nil {
			tx.Rollback()
			return graph, errors.New("Error Occurred!")
		}

		if err := tx.Where("ticket_user_id IN ?", ticket_users).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id).Error; err != nil {
			tx.Rollback()
			return graph, errors.New("Error Occurred!")
		}
	} else {

		if err := tx.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id).Error; err != nil {
			tx.Rollback()
			return graph, errors.New("Error Occurred!")
		}
	}

	db = config.GetDB()
	tx = db.Begin()

	var x time.Time
	var y time.Time
	t := time.Now().Format(DateTime)
	y, _ = time.Parse(DateTime, t)

	graph.TodayDate = time.Now()

	for i := 1; i <= 6; i++ {

		var stats models.TicketStat
		x = y
		y = x.Add(time.Hour * 4)

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("updated_at BETWEEN ?  AND ?", x, y).Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return graph, errors.New("Error Occurred!")
		}

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("created_at BETWEEN ?  AND ?", x, y).Count(&stats.Open)

		switch i {
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
	fmt.Println(weekday)

	y, _ = time.Parse(DateTime, t)
	fmt.Println(y)

	t1 := int(weekday)
	fmt.Println(t1)

	t1 = -t1 + 1

	y = y.AddDate(0, 0, t1)
	graph.StartDate = y

	var count int64

	for i := 1; i <= 7; i++ {

		var stats models.TicketStat
		x = y
		y = x.AddDate(0, 0, 1)

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("updated_at BETWEEN ?  AND ?", x, y).Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return graph, errors.New("Error Occurred!")
		}

		fmt.Println("x", x)

		db.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("created_at BETWEEN ?  AND ?", x, y).Count(&stats.Open)

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
		case 0:
			graph.WeekClosed.Sunday = stats.Closed
			graph.WeekOpen.Sunday = stats.Open
		}

		count += stats.Closed

	}

	graph.EndDate = y
	graph.Sum = count

	return graph, err
}
