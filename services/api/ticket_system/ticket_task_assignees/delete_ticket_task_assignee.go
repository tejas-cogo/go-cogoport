package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketTaskAssignee(id uint) (string,error,uint){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task_assignee models.TicketTaskAssignee

	if err := tx.Model(&ticket_task_assignee).Where("id = ?", id).Update("status","inactive").Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, id
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_task_assignee).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, id
	}

	tx.Commit()
	return "Successfully Created!", err, id
}