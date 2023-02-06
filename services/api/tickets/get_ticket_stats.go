package api

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/constants"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetTicketStats(stats models.TicketStat) (models.TicketStat, error) {
	db := config.GetDB()
	tx := db.Begin()
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


	db = config.GetDB()

	if stats.StartDate != "" {
		start_date, _ := time.Parse(constants.DateTimeFormat(), stats.StartDate)
		end_date, _ := time.Parse(constants.DateTimeFormat(), stats.EndDate)
		end_date = end_date.AddDate(0, 0, 1)

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "unresolved").Where("created_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Unresolved).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("updated_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "rejected").Where("updated_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Rejected).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ? AND tat BETWEEN ? AND ?", "unresolved", t.Format(constants.DateTimeFormat()), time.Now()).Where("created_at BETWEEN ? and ?", start_date, end_date).Count(&stats.DueToday).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "overdue").Where("updated_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Overdue).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "reassigned").Where("created_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Reassigned).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "escalated").Where("created_at BETWEEN ? and ?", start_date, end_date).Count(&stats.Escalated).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

	} else {
		var ticket models.Ticket
		if stats.ExpiryDate != "" || stats.TicketCreatedAt != "" || len(stats.Tags) != 0 || stats.TicketUserID != 0 || stats.QFilter != "" {
			if stats.ExpiryDate != "" {
				ExpiryDate, _ := time.Parse(constants.DateTimeFormat(), stats.ExpiryDate)
				x := ExpiryDate
				y := x.AddDate(0, 0, 1)
				tx = tx.Where("expiry_date BETWEEN ? AND ?", x, y)
			}

			if stats.TicketCreatedAt != "" {
				CreatedAt, _ := time.Parse(constants.DateTimeFormat(), stats.TicketCreatedAt)
				x := CreatedAt
				y := x.AddDate(0, 0, 1)
				tx = tx.Where("created_at BETWEEN ? AND ?", x, y)
			}

			if len(stats.Tags) != 0 {
				tx = tx.Where("tags && ?", "{"+strings.Join(stats.Tags, ",")+"}")
			}

			if stats.TicketUserID != 0 {
				tx = tx.Where("ticket_user_id = ?", stats.TicketUserID)
			}

			if stats.QFilter != "" {
				tx = tx.Where("id::text ilike ? OR type ilike ?", stats.QFilter, "%"+stats.QFilter+"%")
			}
			tx.Where("id IN ?", ticket_id).Distinct("id").Find(&ticket).Pluck("id ", &ticket_id)
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "unresolved").Count(&stats.Unresolved).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "rejected").Count(&stats.Rejected).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ? AND tat BETWEEN ? AND ?", "unresolved", t.Format(constants.DateTimeFormat()), time.Now()).Count(&stats.DueToday).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "overdue").Count(&stats.Overdue).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.TicketActivity{}).Where("id IN ?", ticket_id).Where("status = ?", "reassigned").Count(&stats.Reassigned).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.TicketActivity{}).Where("id IN ?", ticket_id).Where("status = ?", "escalated").Count(&stats.Escalated).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("priority = 'high'").Count(&stats.HighPriority).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		fmt.Println("no stats", ticket_id)

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("expiry_date BETWEEN ? AND ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}
	}

	tx.Commit()
	return stats, err
}
