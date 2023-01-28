package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketDefault(filters models.TicketDefaultType) ([]models.TicketDefault, *gorm.DB) {
	db := config.GetDB()

	var ticket_default []models.TicketDefault

	db = db.Model(&models.TicketDefaultType{})

	db = db.Select("ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status as type_status,ticket_default_timings.id as ticket_default_timing_id,ticket_default_timings.status as timing_status,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.ticket_priority as ticket_priority,ticket_default_timings.status as timing_status ,ticket_default_types.additional_options as additional_options,ticket_default_groups.id as ticket_default_group_id,groups.name as ticket_default_group_name,groups.tags, ticket_default_groups.status as ticket_default_group_status,Count(group_members.id) as member_count")

	db = db.Joins("left join ticket_default_groups on ticket_default_groups.ticket_type = ticket_default_types.ticket_type")

	db = db.Joins("left join ticket_default_timings on ticket_default_timings.ticket_type = ticket_default_types.ticket_type")

	db = db.Joins("left join groups on groups.id = ticket_default_groups.group_id")

	db = db.Joins("left join group_members on groups.id = group_members.group_id")

	db = db.Where("ticket_default_types.ticket_type != ?", "default")

	if filters.TicketType != "" {
		db = db.Where("ticket_default_types.ticket_type = ?", filters.TicketType)
	}

	db = db.Group("ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status,ticket_default_timings.id,ticket_default_timings.status ,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.ticket_priority ,ticket_default_timings.status ,ticket_default_groups.id ,groups.name ,groups.tags, ticket_default_groups.status ")

	db = db.Scan(&ticket_default)

	return ticket_default, db
}
