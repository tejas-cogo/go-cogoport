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
		db.Create(&ticket_user)
		return ticket_user, err
	} else {
		return exist_user, err
	}
	// result := map[string]interface{}{}
}
