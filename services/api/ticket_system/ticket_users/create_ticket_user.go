package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

type TicketUserService struct {
	TicketUser models.TicketUser
}

func CreateTicketUser(ticket_user models.TicketUser) (models.TicketUser,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	ticket_user.Status = "active"
	var exist_user models.TicketUser

	if err := tx.Where("system_user_id = ? and status = ?", ticket_user.SystemUserID, "active").First(&exist_user).Error; err != nil {
		tx.Rollback()
		return ticket_user, errors.New("Error Occurred!")
	}

	tx.Commit()

	if exist_user.ID <= 0 {
		stmt := validate(ticket_user)
		if stmt != "validated" {
			return ticket_user, errors.New(stmt)
		}
		if err := tx.Create(&ticket_user).Error; err != nil {
			tx.Rollback()
			return ticket_user, errors.New("Error Occurred!")
		}
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
