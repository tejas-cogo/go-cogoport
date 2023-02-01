package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
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
		stmt := validations.validate_ticket_user(ticket_user)
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
