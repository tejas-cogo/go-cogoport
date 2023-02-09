package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

func ListTicketDetail(filters models.TicketExtraFilter) (models.TicketDetail, error) {

	var ticket_detail models.TicketDetail
	db := config.GetDB()
	tx := db.Begin()
	var err error
	var ticket models.Ticket
	var ticket_reviewer models.TicketReviewer
	var ticket_reviewer_data models.TicketReviewerData
	// var ticket_spectator models.TicketSpectator

	if err := tx.Where("id = ?", filters.ID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket_detail, errors.New(err.Error())
	}
	ticket_detail.TicketID = ticket.ID
	ticket_detail.Ticket = ticket

	if err := tx.Model(&ticket_reviewer).Where("ticket_id = ? and status != ?", filters.ID, "inactive").Scan(&ticket_reviewer_data).Error; err != nil {
		tx.Rollback()
		return ticket_detail, errors.New(err.Error())
	}
	ticket_detail.TicketReviewerID = ticket_reviewer.ID
	ticket_detail.TicketReviewer = ticket_reviewer_data

	var t_user models.TicketUser
	db.Where("system_user_id = ?", ticket.UserID).First(&t_user)
	ticket_detail.TicketUser = t_user

	var users []string
	users = append(users, ticket_detail.TicketReviewer.UserID.String())
	user_data := helpers.GetUserData(users)
	ticket_detail.TicketReviewer.User.ID = user_data[0].ID
	ticket_detail.TicketReviewer.User.Name = user_data[0].Name
	ticket_detail.TicketReviewer.User.Email = user_data[0].Email
	ticket_detail.TicketReviewer.User.MobileNumber = user_data[0].MobileNumber

	var roles []string
	roles = append(users, ticket_detail.TicketReviewer.RoleID.String())
	role_data := helpers.GetAuthRoleData(roles)
	ticket_detail.TicketReviewer.Role.ID = role_data[0].ID
	ticket_detail.TicketReviewer.Role.Name = role_data[0].Name

	var ticket_default_type models.TicketDefaultType
	var user models.User
	db.Where("ticket_type = ? and status = ?", ticket.Type, "active").First(&ticket_default_type)

	closure_data := helpers.GetUserData(ticket_default_type.ClosureAuthorizers)

	for i := 0; i < len(closure_data); i++ {
		user.ID = closure_data[i].ID
		user.Name = closure_data[i].Name
		user.Email = closure_data[i].Email
		user.MobileNumber = closure_data[i].MobileNumber

		ticket_detail.ClosureAuthorizers = append(ticket_detail.ClosureAuthorizers, user)
	}

	tx.Commit()
	return ticket_detail, err
}
