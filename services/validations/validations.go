package validation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	models "github.com/tejas-cogo/go-cogoport/models"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

func ValidateTicketDefaultRole(ticket_default_role models.TicketDefaultRole) string {

	if ticket_default_role.RoleID == uuid.Nil {
		return ("Role Is Required!")
	}

	if ticket_default_role.Level <= 0 || ticket_default_role.Level > 3 {
		return ("Level Must be Present Between 1 and 3!")
	}

	if ticket_default_role.TicketDefaultTypeID == 0 {
		return ("TicketDefaultTypeID Is Required!")
	}

	return ("validated")
}

func ValidateTicketDefaultTiming(ticket_default_timing models.TicketDefaultTiming) string {

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

func ValidateTicketDefaultType(ticket_default_type models.TicketDefaultType) string {
	if ticket_default_type.TicketType == "" {
		return ("TicketType Is Required!")
	}
	var existed_ticket_default_type models.TicketDefaultType
	db := config.GetDB()
	db.Where("ticket_type = ? and status = ?", ticket_default_type.TicketType, "active").First(&existed_ticket_default_type)

	if existed_ticket_default_type.ID != 0 {
		return ("Ticket type already exists!")
	}

	return ("validated")
}

func ValidateTicketReviewer(ticket_reviewer models.TicketReviewer) string {
	if ticket_reviewer.RoleID == uuid.Nil {
		return ("Role Is Required!")
	}

	if ticket_reviewer.UserID == uuid.Nil {
		return ("User Is Required!")
	}

	// if len(ticket_reviewer.ReviewerManagerIDs) == 0 {
	// 	return ("Invalid Manager Ids Is Required!")
	// }

	if ticket_reviewer.TicketID == 0 {
		return ("Ticket Is Required!")
	}

	return ("validated")
}

func ValidateTokenTicket(ticket models.Ticket) string {
	if ticket.Source == "" {
		return ("Source is Required!")
	}
	if ticket.Type == "" {
		return ("TicketType is Required!")
	}
	if ticket.TicketUserID <= 0 {
		return ("TicketUserID is Required!")
	}
	if ticket.UserType == "" {
		return ("UserType is Required!")
	}

	return ("validated")
}

func ValidateTicketUser(ticket_user models.TicketUser) string {

	var existed_ticket_user models.TicketUser
	if ticket_user.Name == "" {
		return ("User name is Required!")
	}
	if len(ticket_user.Name) < 2 || len(ticket_user.Name) > 40 {
		return ("Name field must be between 2-40 chars!")
	}
	if ticket_user.Email == "" {
		return ("Email is Required!")
	}
	if ticket_user.Type != "client" {
		return ("Type should be client!")
	}

	if ticket_user.Source == "" {
		return ("Source is Required!")
	}

	db := config.GetDB()
	db.Where("name = ?", ticket_user.Name).First(&existed_ticket_user)

	if ticket_user.ID != 0 {
		return ("Name already exists!")
	}

	return ("validated")
}

func ValidateTicket(ticket models.Ticket) string {
	if ticket.Type == "" {
		return ("Ticket Type Is Required!")
	}
	if ticket.Tat == time.Now() {
		return ("Tat couldn't be set!")
	}
	if ticket.ExpiryDate == time.Now() {
		return ("Expiry Date  couldn't be set!")
	}

	return ("validated")
}

func ValidateTicketActivity(ticket_activity models.TicketActivity) string {
	db := config.GetDB()
	var ticket models.Ticket

	if ticket_activity.Status != "assigned" {
		db.Where("id = ?", ticket_activity.TicketID).First(&ticket)
		if ticket.Status != "unresolved" {
			return ("Ticket is not open for activities anymore!")
		}
	}

	if ticket_activity.Status == "" {
		return ("Status is Required!")
	}
	if ticket_activity.TicketID <= 0 {
		return ("TicketID is Required!")
	}
	if ticket_activity.UserID == uuid.Nil {
		return ("UserID is Required!")
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

func ValidateActivityPermission(ticket_activity models.TicketActivity) bool {
	db := config.GetDB()

	var ticket_reviewer models.TicketReviewer
	var ticket models.Ticket

	fmt.Println("ticket_activity", ticket_activity)

	db.Where("ticket_id = ? and status = ?", ticket_activity.TicketID, "active").First(&ticket_reviewer)

	db.Where("id = ? and status = ?", ticket_activity.TicketID, "unresolved").First(&ticket)

	fmt.Println("ticket")
	if ticket_reviewer.UserID != ticket_activity.UserID && ticket.UserID != ticket_activity.UserID {
		return false
	}
	return true
}

func ValidateDuplicateDefaultType(ticket_default_role models.TicketDefaultRole) string{
	var user_ids []string
	var role_ids []string

	db := config.GetDB()
	db.Where("ticket_default_type_id = ? and status = ?", ticket_default_role.TicketDefaultTypeID, "active").Pluck("user_id", &user_ids)
	db.Where("ticket_default_type_id = ? and status = ?", ticket_default_role.TicketDefaultTypeID, "active").Pluck("role_id", &role_ids)

	if !(helpers.Inslice(ticket_default_role.UserID.String(), user_ids)) {
		return ("Cannot assign the user again for this type!")
	}

	if !(helpers.Inslice(ticket_default_role.UserID.String(), role_ids)) {
		return ("Cannot assign the role again for this type!")
	}

	return ("validated")
}