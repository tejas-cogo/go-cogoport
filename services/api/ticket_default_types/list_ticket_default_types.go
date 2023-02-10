package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketDefaultType(filters models.TicketDefaultFilter) ([]models.TicketDefault, *gorm.DB) {
	db := config.GetDB()
	var ticket_default []models.TicketDefault

	db = db.Model(&models.TicketDefaultType{})

	db = db.Select("ticket_type, ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status as type_status,ticket_default_types.closure_authorizer as closure_authorizer,ticket_default_timings.id as ticket_default_timing_id,ticket_default_timings.status as timing_status,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.ticket_priority,ticket_default_types.additional_options")

	db = db.Joins("left join ticket_default_timings on ticket_default_timings.ticket_default_type_id = ticket_default_types.id and ticket_default_timings.deleted_at is null")

	if filters.QFilter != "" {
		filters.QFilter = "%" + filters.QFilter + "%"
		db = db.Where("ticket_default_types.ticket_type iLike ?", filters.QFilter)

	}

	db.Order("ticket_default_types.created desc")

	return ticket_default, db
}
