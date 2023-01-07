package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketReviewer(id uint) uint{
	db := config.GetDB()

	var ticket_reviewer models.TicketReviewer

	db.Model(&ticket_reviewer).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_reviewer)

	return id
}