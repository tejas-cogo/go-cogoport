package validations

import (
	"time"

	models "github.com/tejas-cogo/go-cogoport/models"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

func validate_group_member(group_member models.GroupMember) string {

	if group_member.HierarchyLevel == 0 {
		return ("Hierarchy Level Is Required!")
	}
	if group_member.GroupID == 0 {
		return ("Group Is Required!")
	}

	return ("validated")
}

func validate_group(group models.Group) string {
	if group.Name == "" {
		return ("Group Name Is Required!")
	}

	if len(group.Name) < 2 || len(group.Name) > 40 {
		return ("Name field must be between 2-40 chars!")
	}

	return ("validated")
}

func validate_role(role models.Role) string {
	if role.Name == "" {
		return ("Role Name Is Required!")
	}

	if role.Level == 0 {
		return ("Level Is Required!")
	}

	if role.Level > 9 {
		return ("Level should be in range 1-9!")
	}

	if len(role.Name) < 2 || len(role.Name) > 30 {
		return ("Role field must be between 2-30 chars!")
	}

	return ("validated")
}

func validate_ticket_default_group(ticket_default_group models.TicketDefaultGroup) string {

	if ticket_default_group.GroupID == 0 {
		return ("Group Is Required!")
	}

	if ticket_default_group.TicketDefaultTypeID == 0 {
		return ("TicketDefaultTypeID Is Required!")
	}

	return ("validated")
}

func validate_ticket_default_timing(ticket_default_timing models.TicketDefaultTiming) string {

	if ticket_default_timing.TicketPriority == "" {
		return ("Ticket Priority Is Required!")
	}

	if ticket_default_timing.ExpiryDuration == "" {
		return ("Expiry Duration Is Required!")
	}

	if ticket_default_timing.Tat == "" {
		return ("Tat Is Required!")
	}

	if ticket_default_timing.TicketDefaultTypeID == 0 {
		return ("Ticket Default Type Is Required!")
	}

	ExpiryDuration := helpers.GetDuration(ticket_default_timing.ExpiryDuration)
	Tat := helpers.GetDuration(ticket_default_timing.Tat)

	if ExpiryDuration <= Tat {
		return ("Expiry Duration should be greater than Tat!")
	}

	return ("validated")
}

func validate_ticket_default_type(ticket_default_type models.TicketDefaultType) string {
	if ticket_default_type.TicketType == "" {
		return ("TicketType Is Required!")
	}

	return ("validated")
}

func validate_ticket_reviewer(ticket_reviewer models.TicketReviewer) string {
	if ticket_reviewer.GroupMemberID == 0 {
		return ("Group Member Is Required!")
	}

	if ticket_reviewer.GroupID == 0 {
		return ("Group Is Required!")
	}

	if ticket_reviewer.TicketID == 0 {
		return ("Ticket Is Required!")
	}

	if ticket_reviewer.TicketUserID == 0 {
		return ("Ticket User Is Required!")
	}

	return ("validated")
}

func validate_token_ticket(ticket models.Ticket) string {
	if ticket.Source == "" {
		return ("Source is Required!")
	}
	if ticket.Type != "client" {
		return ("Type should be client!")
	}
	if ticket.TicketUserID <= 0 {
		return ("TicketUserID is Required!")
	}

	return ("validated")
}

func validate_ticket_user(ticket_user models.TicketUser) string {
	if ticket_user.Name == "" {
		return ("User name is Required!")
	}
	if ticket_user.Email == "" {
		return ("Email is Required!")
	}
	if ticket_user.Type != "client" {
		return ("Type should be client!")
	}
	if ticket_user.RoleID != 1 {
		return ("RoleID should be 1!")
	}
	if ticket_user.Source == "" {
		return ("Source is Required!")
	}

	return ("validated")
}

func validate_ticket(ticket models.Ticket) string {
	if ticket.Type == "" {
		return ("Ticket Type Is Required!")
	}
	if ticket.Tat == "" {
		return ("Tat couldn't be set!")
	}
	if ticket.ExpiryDate == time.Now() {
		return ("Expiry Date  couldn't be set!")
	}

	return ("validated")
}

func validate_ticket_activity(ticket_activity models.TicketActivity) string {
	if ticket_activity.Status == "" {
		return ("Status is Required!")
	}
	if ticket_activity.TicketID <= 0 {
		return ("TicketID is Required!")
	}
	if ticket_activity.TicketUserID <= 0 {
		return ("TicketUserID is Required!")
	}
	if ticket_activity.UserType == "" {
		return ("UserType is Required!")
	}
	if ticket_activity.Status == "activity" {
		if ticket_activity.Description == "" {
			return ("Description is Required!")
		}
	}

	return ("validated")
}
