package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
	user "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
	"gorm.io/gorm"
	// tickets "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

type TicketActivityService struct {
	TicketActivity models.TicketActivity
}

func CreateTicketActivity(body models.Filter) (models.TicketActivity, string, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error
	// result := map[string]interface{}{}

	var ticket_user models.TicketUserFilter

	if body.TicketActivity.UserType == "" {
		if body.TicketActivity.TicketUserID == 0 {
			ticket_user.SystemUserID = body.TicketUserFilter.SystemUserID
		} else {
			ticket_user.ID = body.TicketUserFilter.ID
		}

		ticket_user, _ := user.ListTicketUser(ticket_user)
		for _, u := range ticket_user {
			fmt.Println("Fdv", u.ID, "vs")
			body.TicketActivity.UserType = u.Type
			body.TicketActivity.TicketUserID = u.ID
			break
		}
	}
	ticket_activity := body.TicketActivity

	if ticket_activity.Status == "resolved" {
		for _, u := range body.Activity.TicketID {
			var ticket models.Ticket
			ticket_activity.TicketID = u

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Ticket not found", err
			}

			ticket.Status = "closed"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Ticket couldn't be saved", err
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Activity couldn't be created", err
			}
		}
	} else if ticket_activity.Status == "rejected" {
		for _, u := range body.Activity.TicketID {
			var ticket models.Ticket
			ticket_activity.TicketID = u

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "ticket couldn't be find", err
			}
			ticket.Status = "rejected"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Ticket couldn't be saved", err
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Activity couldn't be created", err
			}
		}
	} else if ticket_activity.Status == "escalated" {
		for _, u := range body.Activity.TicketID {
			var group_head models.GroupMember
			ticket_activity.TicketID = u
			var ticket_reviewer models.TicketReviewer
			var ticket models.Ticket

			group_member, err := DeactivateReviewer(u, tx)
			if err != nil {
				tx.Rollback()
				return ticket_activity, "Escalated Group Member not found", err
			}

			if err = tx.Where("group_id = ? and status = ? and hierarchy_level = ?", group_member.GroupID, "active", (group_member.HierarchyLevel)+1).Order("active_ticket_count asc").First(&group_head).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Escalated Group Member not found", err
			}

			group_head.ActiveTicketCount += 1

			if err = tx.Save(&group_head).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Escalated Group Member couldn't be saved", err
			}

			ticket_reviewer.TicketID = u
			ticket_reviewer.TicketUserID = group_head.TicketUserID
			ticket_reviewer.GroupID = group_head.GroupID
			ticket_reviewer.GroupMemberID = group_head.ID
			ticket_reviewer.Status = "active"

			if err = tx.Create(&ticket_reviewer).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Reviewer couldn't be created", err
			}

			ticket.Status = "escalated"
			audits.CreateAuditTicket(ticket, tx)

			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Activity couldn't be created", err
			}
		}
	} else if ticket_activity.Status == "activity" {
		for _, u := range body.Activity.TicketID {
			var ticket models.Ticket
			ticket_activity.TicketID = u
			audits.CreateAuditTicket(ticket, tx)
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, "Activity couldn't be created", err
			}
		}
	} else {
		var ticket models.Ticket
		audits.CreateAuditTicket(ticket, tx)
		if err = tx.Create(&ticket_activity).Error; err != nil {
			tx.Rollback()
			return ticket_activity, "Activity couldn't be created", err
		}
	}

	tx.Commit()

	return ticket_activity, "Successfully Created!", err
}

func DeactivateReviewer(ID uint, tx *gorm.DB) (models.GroupMember, error) {
	var ticket_reviewer models.TicketReviewer
	var group_member models.GroupMember
	var err error

	if err := tx.Where("ticket_id = ? and status = ?", ID, "active").First(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return group_member, err
	}

	ticket_reviewer.Status = "inactive"

	if err := tx.Save(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return group_member, err
	}

	fmt.Println("ticket_reviewer", ticket_reviewer)

	if err := tx.Where("ticket_user_id = ? and status = ?", ticket_reviewer.TicketUserID, "active").First(&group_member).Error; err != nil {
		tx.Rollback()
		return group_member, err
	}

	fmt.Println("group_member", &group_member)
	group_member.ActiveTicketCount = group_member.ActiveTicketCount - 1

	if err := tx.Save(&group_member).Error; err != nil {
		tx.Rollback()
		return group_member, err
	}

	return group_member, err
}
