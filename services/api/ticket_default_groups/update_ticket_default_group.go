package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func UpdateTicketDefaultGroup(body models.TicketDefaultGroup) (models.TicketDefaultGroup,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_group models.TicketDefaultGroup

	if err := tx.Where("id = ?", body.ID).Find(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return body, errors.New(err.Error())
	}

	if body.TicketDefaultTypeID > 0 {
		ticket_default_group.TicketDefaultTypeID = body.TicketDefaultTypeID
	}
	if body.GroupID > 0 {
		ticket_default_group.GroupID = body.GroupID
	}
	if body.GroupMemberID >= 0 {
		ticket_default_group.GroupMemberID = body.GroupMemberID
	}
	if body.Status != "" {
		ticket_default_group.Status = body.Status
	}

	if err := tx.Save(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return body, errors.New(err.Error())
	}

	tx.Commit()
	
	return ticket_default_group, err
}
