package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketDefaultType(filters models.TicketDefaultFilter) ([]models.TicketDefault, *gorm.DB) {
	db := config.GetDB()
	var ticket_default []models.TicketDefault

	ticket_default_group_query := db.Model(&models.TicketDefaultGroup{}).Select("ticket_default_groups.id as ticket_default_group_id,ticket_default_groups.level as group_level,ticket_default_groups.status as status ,ticket_default_groups.ticket_default_type_id as ticket_default_type_id,json_agg(groups.*) as groups ,json_agg(group_members.*) as group_members").Joins("left join (?) groups on groups.group_id = ticket_default_groups.group_id", group_query).Joins("left join (?) group_members on group_members.group_member_id = ticket_default_groups.group_member_id", group_member_query).Where("ticket_default_groups.deleted_at is null").Group("ticket_default_groups.id,ticket_default_groups.level,ticket_default_groups.ticket_default_type_id")

	ticket_default_group_type_query := db.Model(&models.TicketDefaultType{}).Select("ticket_default_types.id as ticket_default_type_id,json_agg(default_groups.*) ticket_default_groups").Joins("left join (?) default_groups on default_groups.ticket_default_type_id = ticket_default_types.id", ticket_default_group_query).Group("ticket_default_types.id")

	db = db.Model(&models.TicketDefaultType{})

	db = db.Select("ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status as type_status,ticket_default_timings.id as ticket_default_timing_id,ticket_default_timings.status as timing_status,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.ticket_priority as ticket_priority,ticket_default_types.additional_options as additional_options,json_agg(ticket_default_groups.*)ticket_default_groups")

	db = db.Joins("left join (?) ticket_default_groups on ticket_default_groups.ticket_default_type_id = ticket_default_types.id ", ticket_default_group_type_query)

	db = db.Joins("left join ticket_default_timings on ticket_default_timings.ticket_default_type_id = ticket_default_types.id and ticket_default_timings.deleted_at is null")

	if filters.QFilter != "" {
		filters.QFilter = "%" + filters.QFilter + "%"
		db = db.Where("ticket_default_types.ticket_type iLike ?", filters.QFilter)

	}

	db = db.Where("ticket_default_types.deleted_at is null")

	db = db.Group("ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status,ticket_default_timings.id,ticket_default_timings.status ,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.ticket_priority")

	return ticket_default, db
}
