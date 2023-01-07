package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func ListTicket(filters models.Ticket, tags string) []models.Ticket {
	db := config.GetDB()

	var ticket []models.Ticket
	result := map[string]interface{}{}

	if (filters.Type != ""){
		db = db.Where("type = ?", filters.Type)
	} 

	if (filters.Priority != ""){
		db = db.Where("priority = ?", filters.Priority)
	} 

	if (tags != ""){
		db = db.Where("? Like ANY(tags)", tags)
	} 

	if (filters.Status != ""){
		db = db.Where("status = ?", filters.Status)
	}else{
		db = db.Where("status = ?", "active")
	} 

	db.Find(&ticket).Take(&result)

	return ticket
}
