package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func UpdateTicketUser(body models.TicketUserRole) ([]models.TicketUser, error) {
	db := config.GetDB()
	var ticket_user []models.TicketUser
	tx := db.Begin()
	var err error

	if err := tx.Where("id IN ?", body.ID).Find(&ticket_user).Error; err != nil {
		tx.Rollback()
		return ticket_user, errors.New("System User Not Found")
	}

	for _, u := range ticket_user {
		if body.Type != "" {
			u.Type = body.Type
		}

		if body.RoleID > 0 {
			u.RoleID = body.RoleID
		}
		if body.Source != "" {
			u.Source = body.Source
		}

		if err := tx.Save(&u).Error; err != nil {
			tx.Rollback()
			return ticket_user, errors.New("System User Not Found")
		}
	}

	tx.Commit()
	return ticket_user, err
}
