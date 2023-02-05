package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	// validations "github.com/tejas-cogo/go-cogoport/services/validations"
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

		// stmt := validations.ValidateTicketUser(ticket_user)

		// if stmt != "validated" {
		// 	return ticket_user, errors.New(stmt)
		// }
		tx := db.Begin()
		if err := tx.Create(&ticket_user).Error; err != nil {
			tx.Rollback()
			return ticket_user, errors.New(err.Error())
		}
		tx.Commit()
		return ticket_user, err

	} else {

		return exist_user, err
	}
}
