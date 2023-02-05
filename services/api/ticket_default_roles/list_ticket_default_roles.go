package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketDefaultRole(filters models.TicketDefaultRole) ([]models.TicketDefaultRole, error) {
	db := config.GetDB()

	var err error

	var ticket_default_roles []models.TicketDefaultRole

	if filters.TicketDefaultTypeID > 0 {
		db = db.Where("ticket_default_type_id = ?", filters.TicketDefaultTypeID)
	}

	if filters.RoleID != 0 {
		db = db.Where("role_id = ?", filters.RoleID)
	}

	if filters.UserID != 0 {
		db = db.Where("user_id = ?", filters.UserID)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Order("created_at desc").Find(&ticket_default_roles)

	db = db.Order("created_at desc").Find(&ticket_default_roles)

	return ticket_default_roles, err
}
