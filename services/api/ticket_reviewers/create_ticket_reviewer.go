package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	groupmember "github.com/tejas-cogo/go-cogoport/services/api/group_members"
	activity "github.com/tejas-cogo/go-cogoport/services/api/ticket_activities"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type TicketReviewerService struct {
	TicketReviewer   models.TicketReviewer
	ReviewerActivity models.ReviewerActivity
}

func CreateTicketReviewer(body models.Ticket) (models.Ticket, error) {
	db := config.GetDB()

	txt := db.Begin()

	var filters models.Filter
	var ticket_reviewer models.TicketReviewer
	var ticket_default_type models.TicketDefaultType
	var ticket_default_group models.TicketDefaultGroup
	var err error

	if err := txt.Where("ticket_type = ? and status = ?", body.Type, "active").First(&ticket_default_type).Error; err != nil {
		if err := txt.Where("id = ?", 1).First(&ticket_default_type).Error; err != nil {
			txt.Rollback()
			return body, errors.New(err.Error())
		}
	}

	if erro := txt.Where("ticket_default_type_id = ? and status = ?", ticket_default_type.ID, "active").First(&ticket_default_group).Error; erro != nil {
		if err := txt.Where("ticket_default_type_id = ? ", 1).First(&ticket_default_group).Error; err != nil {
			txt.Rollback()
			return body, errors.New(err.Error())
		}
	}

	ticket_reviewer.GroupID = ticket_default_group.GroupID
	var group_member models.GroupMember
	if ticket_default_group.GroupMemberID < 1 {
		txt.Where("group_id = ? and status = ? and ticket_user_id != ?", ticket_default_group.GroupID, "active", body.TicketUserID).Order("active_ticket_count asc").Order("hierarchy_level desc").First(&group_member)
	} else {
		txt.Where("id = ?", ticket_default_group.GroupMemberID).First(&group_member)
	}

	ticket_reviewer.GroupMemberID = group_member.ID
	ticket_reviewer.TicketUserID = group_member.TicketUserID
	ticket_reviewer.TicketID = body.ID
	filters.GroupMember.ID = group_member.ID
	ticket_reviewer.Status = "active"

	stmt := validations.ValidateTicketReviewer(ticket_reviewer)
	if stmt != "validated" {
		return body, errors.New(stmt)
	}
	if err := txt.Create(&ticket_reviewer).Error; err != nil {
		txt.Rollback()
		return body, errors.New(err.Error())
	}

	filters.GroupMember.ActiveTicketCount = group_member.ActiveTicketCount + 1
	groupmember.UpdateGroupMember(filters.GroupMember)

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = ticket_reviewer.TicketID
	ticket_activity.TicketUserID = ticket_reviewer.TicketUserID
	ticket_activity.UserType = "system"
	ticket_activity.Type = "reviewer_assigned"
	ticket_activity.Status = "assigned"

	stmt3 := validations.ValidateTicketActivity(ticket_activity)
	if stmt3 != "validated" {
		return body, errors.New(stmt)
	}
	if err := txt.Create(&ticket_activity).Error; err != nil {
		txt.Rollback()
		return body, errors.New(err.Error())
	}

	if ticket_activity.UserType == "internal" {
		activity.SendTicketActivity(ticket_activity)
	}

	txt.Commit()
	return body, err
}
