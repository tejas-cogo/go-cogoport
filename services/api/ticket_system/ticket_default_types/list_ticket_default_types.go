package ticket_system

import (
	"fmt"
	"errors"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketDefaultType(filters models.TicketDefaultFilter) ([]models.TicketDefault, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default []models.TicketDefault

	tx = tx.Model(&models.TicketDefaultType{})

	tx = tx.Select("ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status as type_status,ticket_default_timings.id as ticket_default_timing_id,ticket_default_timings.status as timing_status,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.ticket_priority as ticket_priority,ticket_default_timings.status as timing_status ,ticket_default_types.additional_options as additional_options,ticket_default_groups.id as ticket_default_group_id,groups.name as group_name,groups.id as group_id,groups.tags,ticket_default_groups.group_member_id as group_member_id,ticket_users.name as group_member_name ,ticket_default_groups.status as ticket_default_group_status,Count(group_members.id) as member_count,ticket_users.email as group_member_email")

	tx = tx.Joins("left join ticket_default_groups on ticket_default_groups.ticket_default_type_id = ticket_default_types.id")

	tx = tx.Joins("left join group_members gpm on ticket_default_groups.group_member_id = gpm.id ")

	tx = tx.Joins("left join ticket_users on gpm.ticket_user_id = ticket_users.id")

	tx = tx.Joins("left join ticket_default_timings on ticket_default_timings.ticket_default_type_id = ticket_default_types.id")

	tx = tx.Joins("left join groups on groups.id = ticket_default_groups.group_id")

	tx = tx.Joins("left join group_members on groups.id = group_members.group_id")

	tx = tx.Where("ticket_default_types.id != ?", 1)

	fmt.Println("ewdfs", filters.QFilter, "Q")
	if filters.QFilter != "" {
		filters.QFilter = "%" + filters.QFilter + "%"
		tx = tx.Where("ticket_default_types.ticket_type iLike ?", filters.QFilter)

	}

	tx = tx.Group("ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status,ticket_default_timings.id,ticket_default_timings.status ,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.ticket_priority ,ticket_default_timings.status ,ticket_default_groups.id ,groups.name ,groups.tags, ticket_default_groups.status,ticket_users.name,gpm.id,groups.id,ticket_users.email")
	if err := tx.Error; err != nil {
		tx.Rollback()
		return ticket_default, tx, errors.New("Error Occurred!")
	}

	tx.Commit()
	return ticket_default, tx, err
}
