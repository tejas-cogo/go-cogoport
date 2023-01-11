package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketUser(filters models.TicketUser) []models.TicketUser {
	db := config.GetDB()

	var ticket_user []models.TicketUser

	if (filters.ID != 0){
		db = db.Where("id = ?", filters.ID)
	} 

	// if (filters.Priority != ""){
	// 	db = db.Where("priority = ?", filters.Priority)
	// } 

	// if (filters.Source != ""){
	// 	db = db.Where("source = ?", filters.Source)
	// }
 
	db.Find(&ticket_user)

	return ticket_user
}
