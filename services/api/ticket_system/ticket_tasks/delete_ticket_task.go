package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketTask(id uint) (string,error,uint){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task models.TicketTask

	db.Model(&ticket_task).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_task)

	tx.Commit()
	return "Successfully Created!", err, id
}