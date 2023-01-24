package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketDefaultGroupService struct {
	TicketDefaultGroup models.TicketDefaultGroup
}

func CreateTicketDefaultGroup(ticket_default_group models.TicketDefaultGroup) (string, error) {
	db := config.GetDB()
	// result := map[string]interface{}{}
	tx := db.Begin()
	var err error

	stmt := validate(ticket_default_group)
	if stmt != "validated" {
		return stmt, err
	}

	ticket_default_group.Status = "active"

	if err := tx.Create(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err
	}

	tx.Commit()

	return "Successfully Created!", err
}

func validate(ticket_default_group models.TicketDefaultGroup) string {
	if ticket_default_group.TicketType == "" {
		return ("TicketType Is Required!")
	}

	if ticket_default_group.GroupID == 0 {
		return ("Group Is Required!")
	}

	if ticket_default_group.TicketDefaultTypeID == 0 {
		return ("TicketDefaultType Is Required!")
	}

	return ("validated")
}
