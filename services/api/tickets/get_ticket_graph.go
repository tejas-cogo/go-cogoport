package api

import (
	"errors"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/constants"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetTicketGraph(graph models.TicketGraph) (models.TicketGraph, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_reviewer []models.TicketReviewer
	var ticket_id []uint

	if graph.AgentRmID != "" {

		db.Where("manager_rm_ids && '(?)' or user_id = ?", graph.AgentRmID, graph.AgentRmID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	} else if graph.AgentID != "" {

		db.Where("user_id = ?", graph.AgentID).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	} else {

		db.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	}

	db = config.GetDB()
	tx = db.Begin()

	var x time.Time
	var y time.Time
	t := time.Now().Format(constants.DateTimeFormat())
	y, _ = time.Parse(constants.DateTimeFormat(), t)

	graph.TodayDate = time.Now()

	for i := 1; i <= 6; i++ {

		var stats models.TicketStat
		x = y
		y = x.Add(time.Hour * 4)

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("updated_at BETWEEN ?  AND ?", x, y).Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return graph, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ? or status = ?", "unresolved", "pending").Where("created_at BETWEEN ?  AND ?", x, y).Count(&stats.Open).Error; err != nil {
			tx.Rollback()
			return graph, errors.New(err.Error())
		}

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

	y, _ = time.Parse(constants.DateTimeFormat(), t)

	t1 := int(weekday)

	if t1 == 0 {
		t1 = -6
	} else {
		t1 = -t1 + 1
	}

	y = y.AddDate(0, 0, t1)
	graph.StartDate = y

	var count int64

	for i := 1; i <= 7; i++ {

		var stats models.TicketStat
		x = y
		y = x.AddDate(0, 0, 1)

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("updated_at BETWEEN ? AND ?", x, y).Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return graph, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("created_at BETWEEN ? AND ?", x, y).Where("status = ? or status = ?", "unresolved", "pending").Count(&stats.Open).Error; err != nil {
			tx.Rollback()
			return graph, errors.New(err.Error())
		}

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

	tx.Commit()
	return graph, err
}
