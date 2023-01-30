package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketDefaultGroup(body models.TicketDefaultGroup) (string,error,models.TicketDefaultGroup) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_group models.TicketDefaultGroup

	if err := tx.Where("id = ?", body.ID).Find(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	if body.TicketDefaultTypeID > 0 {
		ticket_default_group.TicketDefaultTypeID = body.TicketDefaultTypeID
	}
	if body.GroupID != 0 {
		ticket_default_group.GroupID = body.GroupID
	}
	if body.GroupMemberID != 0 {
		ticket_default_group.GroupMemberID = body.GroupMemberID
	}
	if body.Status != "" {
		ticket_default_group.Status = body.Status
	}

	if err := tx.Save(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	tx.Commit()
	
	return "Sucessfully Updated!", err, ticket_default_group
}
