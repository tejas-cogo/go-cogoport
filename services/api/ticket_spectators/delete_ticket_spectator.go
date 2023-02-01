package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_activities"
	"errors"
)

func DeleteTicketSpectator(id uint) (uint,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_spectator models.TicketSpectator
	var spectator_activity models.SpectatorActivity
	var filters models.Filter

	if err := tx.Model(&ticket_spectator).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot update ticket spectator status!")
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_spectator).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot delete ticket spectator!")
	}

	spectator_activity.TicketID = ticket_spectator.TicketID
	spectator_activity.PerformedByID = ticket_spectator.PerformedByID
	filters.TicketActivity.Type = "Spectator Deleted"
	filters.TicketActivity.Status = "deleted"

	activities.CreateTicketActivity(filters)

	tx.Commit()
	return id, err
}
