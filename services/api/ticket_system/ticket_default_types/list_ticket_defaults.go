package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketDefault(filters models.Filter) ([]models.TicketDefault, *gorm.DB) {

	db := config.GetDB()

	// ticket_type_data, db := ListTicketDefaultType(filters.TicketDefaultType)

	var ticket_default []models.TicketDefault
	// var ticket_defaults []interface{}

	db = db.Model(&models.TicketDefaultType{}).Select("ticket_default_types.id, ticket_default_types.ticket_type,ticket_default_types.status as type_status,ticket_default_timings.id as ticket_default_timing_id,ticket_default_timings.status as timing_status,ticket_default_timings.expiry_duration ,ticket_default_timings.tat ,ticket_default_timings.conditions,ticket_default_timings.status as timing_status ,ticket_default_groups.id as ticket_default_group_id,groups.name as ticket_default_group_name,groups.tags, ticket_default_groups.status as ticket_default_group_status").Joins("left join ticket_default_groups on ticket_default_groups.ticket_type = ticket_default_types.ticket_type").Joins("left join ticket_default_timings on ticket_default_timings.ticket_type = ticket_default_types.ticket_type").Joins("left join groups on groups.id = ticket_default_groups.group_id").Scan(&ticket_default)

	// for _, u := range ticket_type_data {

	// 	ticket_default.TicketDefaultType = u

	// 	filters.TicketDefaultGroup.TicketType = u.TicketType
	// 	ticket_group_data, _ := ticketdefaultgroup.ListTicketDefaultGroup(filters.TicketDefaultGroup)
	// 	for _, v := range ticket_group_data {
	// 		ticket_default.TicketDefaultGroup = v
	// 		break
	// 	}

	// 	filters.TicketDefaultTiming.TicketType = u.TicketType
	// 	ticket_timing_data, _ := ticketdefaulttiming.ListTicketDefaultTiming(filters.TicketDefaultTiming)
	// 	for _, x := range ticket_timing_data {
	// 		ticket_default.TicketDefaultTiming = x
	// 		break
	// 	}
	// 	log.Println(ticket_default)

	// 	ticket_defaults = append(ticket_defaults, )

	// }
	fmt.Println("s", ticket_default, "s")
	return ticket_default, db
}
