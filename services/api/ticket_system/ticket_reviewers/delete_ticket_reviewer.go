package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketReviewer(id uint) (string,error,uint) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_reviewer models.TicketReviewer
	var group_member models.GroupMember

	if err := tx.Model(&ticket_reviewer).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err ,id
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err ,id
	}

	if err := tx.Where("id = ?", ticket_reviewer.GroupMemberID).First(&group_member).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err ,id
	}

	group_member.ActiveTicketCount -= 1

	if err := tx.Save(&group_member).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err ,id
	}

	tx.Commit()

	return "Successfully Deleted!", err, id
}
