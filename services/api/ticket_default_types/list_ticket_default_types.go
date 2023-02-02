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

	db = db.Select("ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status as type_status,ticket_default_timings.id as ticket_default_timing_id,ticket_default_timings.status as timing_status,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.ticket_priority as ticket_priority,ticket_default_timings.status as timing_status ,ticket_default_types.additional_options as additional_options,ticket_default_groups.id as ticket_default_group_id,groups.name as group_name,groups.id as group_id,groups.tags,ticket_default_groups.group_member_id as group_member_id,ticket_users.name as group_member_name ,ticket_default_groups.status as ticket_default_group_status,Count(group_members.id) as member_count,ticket_users.email as group_member_email")

	db = db.Joins("left join ticket_default_groups on ticket_default_groups.ticket_default_type_id = ticket_default_types.id and ticket_default_groups.deleted_at is null")

	db = db.Joins("left join group_members gpm on ticket_default_groups.group_member_id = gpm.id and gpm.deleted_at is null ")

	db = db.Joins("left join ticket_users on gpm.ticket_user_id = ticket_users.id and ticket_users.deleted_at is null")

	db = db.Joins("left join ticket_default_timings on ticket_default_timings.ticket_default_type_id = ticket_default_types.id and ticket_default_timings.deleted_at is null")

	db = db.Joins("left join groups on groups.id = ticket_default_groups.group_id and groups.deleted_at is null")

	db = db.Joins("left join group_members on groups.id = group_members.group_id and group_members.deleted_at is null")

	if filters.QFilter != "" {
		filters.QFilter = "%" + filters.QFilter + "%"
		db = db.Where("ticket_default_types.ticket_type iLike ?", filters.QFilter)

	}

	db = db.Where("ticket_default_types.deleted_at is null")

	db = db.Group("ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status,ticket_default_timings.id,ticket_default_timings.status ,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.ticket_priority ,ticket_default_timings.status ,ticket_default_groups.id ,groups.name ,groups.tags, ticket_default_groups.status,ticket_users.name,gpm.id,groups.id,ticket_users.email")

	return ticket_default, db
}
