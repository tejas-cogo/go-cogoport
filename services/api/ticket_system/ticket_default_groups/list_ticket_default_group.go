package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	// "fmt"
)

func ListTicketDefaultGroup(filters models.TicketDefaultGroup) []models.TicketDefaultGroup{
	db := config.GetDB()

	var ticket_default_groups []models.TicketDefaultGroup

	result := map[string]interface{}{}

	if (filters.TicketType != ""){
		db = db.Where("ticket_type = ?", filters.TicketType)
	} 

	if (filters.GroupID != 0){
		db = db.Where("group_id = ?", filters.GroupID)
	} 

	if (filters.Status != ""){
		db = db.Where("status = ?", filters.Status)
	}else{
		db = db.Where("status = ?", "active")
	} 

	db.Take(&result)

	return ticket_default_groups
}