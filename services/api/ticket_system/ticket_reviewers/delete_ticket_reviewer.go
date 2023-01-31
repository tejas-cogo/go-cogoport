package ticket_system

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketReviewer(id uint) (uint,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_reviewer models.TicketReviewer
	var group_member models.GroupMember

	if err := tx.Model(&ticket_reviewer).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occured!")
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occured!")
	}

	if err := tx.Where("id = ?", ticket_reviewer.GroupMemberID).First(&group_member).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occured!")
	}

	group_member.ActiveTicketCount -= 1

	if err := tx.Save(&group_member).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occured!")
	}

	tx.Commit()

	return id, err
}
