package ticket_system

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketWorkerService struct {
	TicketEscalatedPayload models.TicketEscalatedPayload
}

func GetDuration(ExpiryDuration string) int {
	duration := strings.Split(ExpiryDuration, ":")

	durationd := strings.Split(duration[0], "d")
	durationh := strings.Split(duration[1], "h")
	durationm := strings.Split(duration[2], "m")

	d, _ := strconv.Atoi(durationd[0])
	h, _ := strconv.Atoi(durationh[0])
	m, _ := strconv.Atoi(durationm[0])

	h += m / 60
	h += d * 24

	return h

}

func TicketEscalation(p models.TicketEscalatedPayload) error {

	db := config.GetDB()

	var ticket models.Ticket
	var ticket_reviewer models.TicketReviewer
	var ticket_reviewer_new models.TicketReviewer
	var ticket_default_timing models.TicketDefaultTiming
	var group_member models.GroupMember
	var group_head models.GroupMember

	tx := db.Begin()

	if err := db.Where("id = ?", p.TicketID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return err
	}

	if ticket.Status == "unresolved" {

		if err := db.Where("ticket_type = ?", ticket.Type).First(&ticket_default_timing).Error; err != nil {
			tx.Rollback()
			return err
		}

		ticket.Tat = ticket_default_timing.Tat

		ticket.ExpiryDate = time.Now()
		Duration := GetDuration(ticket_default_timing.ExpiryDuration)
		ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(Duration))

		if err := db.Save(&ticket).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := db.Where("ticket_id = ? and status = 'active'", ticket.ID).First(&ticket_reviewer).Error; err != nil {
			tx.Rollback()
			return err
		}

		ticket_reviewer.Status = "inactive"

		if err := db.Save(&ticket_reviewer).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := db.Where("ticket_user_id = ? and status = 'active'", ticket_reviewer.TicketUserID).Find(&group_member).Error; err != nil {
			tx.Rollback()
			return err
		}

		group_member.ActiveTicketCount -= 1

		if err := db.Save(&group_member).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := db.Where("ticket_user_id = ? and status = 'active'", group_member.GroupHeadID).First(&group_head).Error; err != nil {
			tx.Rollback()
			return err
		}

		ticket_reviewer_new.TicketID = ticket.ID
		ticket_reviewer_new.GroupID = group_head.GroupID
		ticket_reviewer_new.GroupMemberID = group_head.ID
		ticket_reviewer_new.TicketUserID = group_head.TicketUserID

		if err := db.Create(&ticket_reviewer_new).Error; err != nil {
			tx.Rollback()
			return err
		}

		group_head.ActiveTicketCount += 1

		if err := db.Save(&group_head).Error; err != nil {
			tx.Rollback()
			return err
		}

		var ticket_activity models.TicketActivity
		ticket_activity.TicketID = ticket_reviewer.TicketID
		ticket_activity.TicketUserID = ticket_reviewer.TicketUserID
		ticket_activity.UserType = "worker"
		ticket_activity.Type = "Reviewer Escalated"
		ticket_activity.Status = "escalated"

		if err := db.Create(&ticket_activity).Error; err != nil {
			tx.Rollback()
			return err
		}

		var ticket_audit models.TicketAudit

		ticket_audit.ObjectId = ticket.ID
		ticket_audit.Action = "escalation"
		ticket_audit.Object = "Ticket"
		data, _ := json.Marshal(ticket)
		ticket_audit.Data = string(data)

		if err := db.Create(&ticket_audit).Error; err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()

	}
	return tx.Commit().Error
}
