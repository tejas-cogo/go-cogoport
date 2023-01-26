package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketUser(body models.TicketUserRole) (models.TicketUser, string, error) {
	db := config.GetDB()
	var ticket_user models.TicketUser
	tx := db.Begin()
	var err error
	for _, u := range body.ID {

		if err := tx.Where("id = ?", u).First(&ticket_user).Error; err != nil {
			tx.Rollback()
			return ticket_user, "System User Not Found", err
		}

		if body.Type != ticket_user.Type {
			ticket_user.Type = body.Type
		}

		if body.RoleID != ticket_user.RoleID {
			ticket_user.RoleID = body.RoleID
		}
		if body.Source != ticket_user.Source {
			ticket_user.Source = body.Source
		}

		if err := tx.Save(&ticket_user).Error; err != nil {
			tx.Rollback()
			return ticket_user, "System User Not Found", err
		}
	}

	tx.Commit()
	return ticket_user, "Successfully Updated!", err
}
