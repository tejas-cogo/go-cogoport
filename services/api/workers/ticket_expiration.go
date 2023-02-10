package api

import (
	"encoding/json"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func TicketExpiration(p models.TicketPayload) error {

	db := config.GetDB()

	var ticket models.Ticket
	var ticket_reviewer models.TicketReviewer
	var ticket_spectator models.TicketSpectator

	tx := db.Begin()

	if err := tx.Where("id = ? ", p.TicketID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return err
	}

	if ticket.Status == "unresolved" {

		if ticket.ExpiryDate == time.Now() {
			ticket.Status = "overdue"
			if err := tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return err
			}

			if err := tx.Where("ticket_id = ? and status = ?", ticket.ID, "active").First(&ticket_reviewer).Error; err != nil {
				tx.Rollback()
				return err
			}
			ticket_reviewer.Status = "inactive"

			if err := tx.Save(&ticket_reviewer).Error; err != nil {
				tx.Rollback()
				return err
			}

			if err := tx.Where("ticket_id = ? and status = ?", ticket.ID, "active").First(&ticket_spectator).Error; err != nil {
				tx.Rollback()
				return err
			}
			ticket_spectator.Status = "inactive"

			if err := tx.Save(&ticket_spectator).Error; err != nil {
				tx.Rollback()
				return err
			}

			var ticket_activity models.TicketActivity
			ticket_activity.TicketID = ticket_reviewer.TicketID
			ticket_activity.UserID = ticket_reviewer.UserID
			ticket_activity.UserType = "system"
			ticket_activity.Type = "Ticket Expired"
			ticket_activity.Status = "expired"

			if err := tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return err
			}

			var ticket_audit models.TicketAudit

			ticket_audit.ObjectId = ticket.ID
			ticket_audit.Action = "expired"
			ticket_audit.Object = "Ticket"
			data, _ := json.Marshal(ticket)
			ticket_audit.Data = string(data)

			if err := tx.Create(&ticket_audit).Error; err != nil {
				tx.Rollback()
				return err
			}

		}

	}

	tx.Commit()

	return tx.Commit().Error
}
