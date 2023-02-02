package api

import (
	"errors"
	"strings"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

const (
	YYYYMMDD = "2006-01-02"
)

func GetTicketStats(stats models.TicketStat) (models.TicketStat, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_reviewer []models.TicketReviewer
	var ticket_user []models.TicketUser
	var ticket_id []uint
	var ticket_users []uint
	t := time.Now()

	if stats.AgentRmID != "" {

		db2 := config.GetCDB()
		tx2 := db2.Begin()
		var partner_user_rm_mapping []models.PartnerUserRmMapping
		var partner_user_rm_ids []string

		if err := tx2.Where("reporting_manager_id = ? and status = ?", stats.AgentRmID, "active").Distinct("user_id").Find(&partner_user_rm_mapping).Pluck("user_id", &partner_user_rm_ids).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err := tx.Where("system_user_id IN ?", partner_user_rm_ids).Distinct("id").Find(&ticket_user).Pluck("id", &ticket_users).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err := tx.Where("ticket_user_id IN ?", ticket_users).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

	} else if stats.AgentID != "" {
		if err := tx.Where("system_user_id = ?", stats.AgentID).Distinct("id").Find(&ticket_user).Pluck("id", &ticket_users).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err := tx.Where("ticket_user_id IN ?", ticket_users).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}
	} else {

		if err := tx.Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}
	}

	db = config.GetDB()

	if stats.StartDate != "" {
		start_date, _ := time.Parse(YYYYMMDD, stats.StartDate)
		end_date, _ := time.Parse(YYYYMMDD, stats.EndDate)
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

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ? AND tat BETWEEN ? AND ?", "unresolved", t.Format(YYYYMMDD), time.Now()).Where("created_at BETWEEN ? and ?", start_date, end_date).Count(&stats.DueToday).Error; err != nil {
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

	} else if stats.ExpiryDate != "" {
		t, _ := time.Parse(YYYYMMDD, stats.ExpiryDate)
		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("expiry_date BETWEEN ? and ?", t, t.AddDate(0, 0, 1)).Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("expiry_date BETWEEN ? and ?", t, t.AddDate(0, 0, 1)).Where("priority = 'high'").Count(&stats.HighPriority).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("expiry_date BETWEEN ? and ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

	} else if stats.TicketCreatedAt != "" {

		t, _ := time.Parse(YYYYMMDD, stats.TicketCreatedAt)
		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("created_at BETWEEN ? and ?", t, t.AddDate(0, 0, 1)).Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("created_at BETWEEN ? and ?", t, t.AddDate(0, 0, 1)).Where("priority = 'high'").Count(&stats.HighPriority).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("created_at BETWEEN ? and ?", t, t.AddDate(0, 0, 1)).Where("expiry_date BETWEEN ? AND ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

	} else if len(stats.Tags) != 0 {
		if err = tx.Model(&models.Ticket{}).Where("tags && ? and id IN ?", "{"+strings.Join(stats.Tags, ",")+"}", ticket_id).Pluck("id", &ticket_id).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("priority = 'high'").Count(&stats.HighPriority).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("expiry_date BETWEEN ? AND ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

	} else if stats.TicketUserID != 0 {

		if err = tx.Model(&models.TicketReviewer{}).Where("ticket_id IN ? and ticket_user_id = ?", ticket_id, stats.TicketUserID).Pluck("ticket_id", &ticket_id).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("priority = 'high'").Count(&stats.HighPriority).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("expiry_date BETWEEN ? AND ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

	} else if stats.QFilter != "" {

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "closed").Where("id::text ilike ? OR type ilike ?", stats.QFilter, "%"+stats.QFilter+"%").Count(&stats.Closed).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("priority = 'high'").Where("id::text ilike ? OR type ilike ?", stats.QFilter, "%"+stats.QFilter+"%").Count(&stats.HighPriority).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("id::text ilike ? OR type ilike ?", stats.QFilter, "%"+stats.QFilter+"%").Where("expiry_date BETWEEN ? AND ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

	} else {
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

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ? AND tat BETWEEN ? AND ?", "unresolved", t.Format(YYYYMMDD), time.Now()).Count(&stats.DueToday).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status = ?", "overdue").Count(&stats.Overdue).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "reassigned").Count(&stats.Reassigned).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.TicketActivity{}).Where("ticket_id IN ?", ticket_id).Where("status = ?", "escalated").Count(&stats.Escalated).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("priority = 'high'").Count(&stats.HighPriority).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}

		if err = tx.Model(&models.Ticket{}).Where("id IN ?", ticket_id).Where("status != ? and status != ?", "closed", "rejected").Where("expiry_date BETWEEN ? AND ?", t, t.AddDate(0, 0, 1)).Count(&stats.ExpiringSoon).Error; err != nil {
			tx.Rollback()
			return stats, errors.New(err.Error())
		}
	}

	tx.Commit()
	return stats, err
}
