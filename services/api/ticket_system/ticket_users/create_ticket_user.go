package ticket_system

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketUserService struct {
	TicketUser models.TicketUser
}

func CreateTicketUser(ticket_user models.TicketUser) (models.TicketUser, error) {
	db := config.GetDB()

	ticket_user.Status = "active"
	var exist_user models.TicketUser
	var err error
	db.Where("system_user_id = ? and status = ?", ticket_user.SystemUserID, "active").First(&exist_user)

	if exist_user.ID <= 0 {

		ticket_user.RoleID = 1

		stmt := validate(ticket_user)
		if stmt != "validated" {
			return ticket_user, errors.New(stmt)
		}
		tx := db.Begin()
		if err := tx.Create(&ticket_user).Error; err != nil {
			tx.Rollback()
			return ticket_user, errors.New("Error Occurred!")
		}
		tx.Commit()
		return ticket_user, err

	} else {

		return exist_user, err
	}

	// result := map[string]interface{}{}
}

func validate(ticket_user models.TicketUser) string {
	if ticket_user.Name == "" {
		return ("User name is Required!")
	}
	if ticket_user.Email == "" {
		return ("Email is Required!")
	}
	if ticket_user.Type == "" {
		return ("Type is Required!")
	}
	if ticket_user.RoleID != 1 {
		return ("RoleID should be 1!")
	}
	if ticket_user.Source == "" {
		return ("Source is Required!")
	}

	return ("validated")
}
