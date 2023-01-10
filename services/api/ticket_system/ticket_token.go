package token

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	_"crypto/sha256"
)

// type TicketTokenService struct {
// 	TicketTask models.TicketTask
// }

func CreateTicketToken(ticket_task models.TicketTask) models.TicketTask {
	db := config.GetDB()
	// result := map[string]interface{}{}
	//message := []byte(user_id)
	//sha256.Sum256(message)
	db.Create(&ticket_task)
	return ticket_task
}