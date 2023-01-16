package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"

)

func ListTicketDetail(filters models.TicketDetail) (models.TicketDetail, *gorm.DB) {
	db := config.GetDB()

	var ticket_detail models.TicketDetail 
	db.Where("ticket_id = ?",filters.TicketID).Find(&ticket_detail.TicketActivity)
	db.Where("ticket_id = ?",filters.TicketID).Find(&ticket_detail.TicketReviewer)
	db.Where("ticket_id = ?",filters.TicketID).Find(&ticket_detail.TicketSpectator)

	return ticket_detail, db
}
