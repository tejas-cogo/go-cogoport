package api

import (
	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketDefaultRole(filters models.TicketDefaultRole) ([]models.TicketTypeDefaultRole, error) {
	db := config.GetDB()

	var err error

	var ticket_default_roles []models.TicketDefaultRole
	var ticket_default_type_roles []models.TicketTypeDefaultRole
	db = db.Model(&ticket_default_roles)

	if filters.TicketDefaultTypeID > 0 {
		db = db.Where("ticket_default_type_id = ?", filters.TicketDefaultTypeID)
	}

	if filters.RoleID != uuid.Nil {
		db = db.Where("role_id = ?", filters.RoleID)
	}

	if filters.UserID != uuid.Nil {
		db = db.Where("user_id = ?", filters.UserID)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db.Order("created_at desc").Scan(&ticket_default_type_roles)

	return ticket_default_type_roles, err
}
